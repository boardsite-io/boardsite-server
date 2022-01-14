package session

import (
	"context"
	"errors"
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/heat1q/boardsite/api/log"
	"github.com/heat1q/boardsite/redis"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Dispatcher
type Dispatcher interface {
	// GetSCB returns the session control block for given sessionID.
	GetSCB(sessionID string) (Controller, error)
	// Create creates and initializes a new SessionControl struct
	Create(ctx context.Context, maxUsers int) (string, error)
	// Close removes the SCB from the activesession map and closes the session.
	Close(ctx context.Context, sessionID string) error
	// IsValid checks if session with sessionID exists.
	IsValid(sessionID string) bool
}

type sessionsDispatcher struct {
	mu            sync.RWMutex
	activeSession map[string]Controller
	cache         redis.Handler
}

var _ Dispatcher = (*sessionsDispatcher)(nil)

func NewDispatcher(cache redis.Handler) Dispatcher {
	return &sessionsDispatcher{
		activeSession: make(map[string]Controller),
		cache:         cache,
	}
}

func (d *sessionsDispatcher) GetSCB(sessionID string) (Controller, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	scb, ok := d.activeSession[sessionID]
	if !ok {
		return nil, errors.New("session not found")
	}
	return scb, nil
}

func (d *sessionsDispatcher) Create(ctx context.Context, maxUsers int) (string, error) {
	var sid string
	for {
		id, err := gonanoid.Generate(alphabet, 8)
		if err != nil {
			return "", err
		}
		// ensure uniqueness of id
		if _, err := d.GetSCB(id); err != nil {
			sid = id
			break
		}
	}

	scb := NewControlBlock(sid, d.cache, d, maxUsers)
	// assign to SessionControl struct
	d.mu.Lock()
	d.activeSession[scb.id] = scb
	d.mu.Unlock()
	log.Ctx(ctx).Infof("Create Session with ID: %s", scb.id)

	return sid, nil
}

func (d *sessionsDispatcher) Close(ctx context.Context, sessionID string) error {
	scb, err := d.GetSCB(sessionID)
	if err != nil {
		return err
	}
	scb.Close()
	d.mu.Lock()
	delete(d.activeSession, sessionID)
	d.mu.Unlock()

	if err := scb.Attachments().Clear(); err != nil {
		log.Ctx(ctx).Warnf("cannot clear attachment for %s: %v\n", scb.ID(), err)
	}

	log.Ctx(ctx).Infof("Close session %s", scb.ID())

	return nil
}

func (d *sessionsDispatcher) IsValid(sessionID string) bool {
	_, err := d.GetSCB(sessionID)
	return err == nil
}
