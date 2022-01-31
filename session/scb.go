package session

import (
	"context"
	"log"
	"sync"

	"github.com/heat1q/boardsite/api/types"
	"github.com/heat1q/boardsite/attachment"
	"github.com/heat1q/boardsite/redis"
)

type CreateSessionResponse struct {
	SessionId string `json:"sessionId"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Controller
type Controller interface {
	ID() string
	GetPages(ctx context.Context) ([]string, map[string]*PageMeta, error)
	GetPagesSet(ctx context.Context) map[string]struct{}
	IsValidPage(ctx context.Context, pageID ...string) bool
	AddPages(ctx context.Context, pageIDs []string, index []int, meta map[string]*PageMeta) error
	DeletePages(ctx context.Context, pageID ...string) error
	UpdatePages(ctx context.Context, pageIDs []string, meta map[string]*PageMeta, clear bool) error
	SyncPages(ctx context.Context) error
	NewUser(alias string, color string) (*User, error)
	UserReady(u *User) error
	GetUserReady(userID string) (*User, error)
	IsUserReady(userID string) bool
	UserConnect(u *User)
	UserDisconnect(ctx context.Context, userID string)
	IsUserConnected(userID string) bool
	GetUsers() map[string]*User
	Close()
	GetStrokes(ctx context.Context, pageID string) ([]*Stroke, error)
	Receive(ctx context.Context, msg *types.Message) error
	Attachments() attachment.Handler
	// NumUsers returns the number of active users in the session
	NumUsers() int
}

// controlBlock holds the information and channels for sessions
type controlBlock struct {
	id string

	attachments attachment.Handler
	dispatcher  Dispatcher

	broadcast chan *types.Message
	echo      chan *types.Message

	cache    redis.Handler
	dbUpdate chan []redis.Stroke

	signalClose chan struct{}

	muRdyUsr sync.RWMutex
	// users that have previously been created via POST
	// and have not yet joined the session
	usersReady map[string]*User

	muUsr sync.RWMutex
	// Active Client users that are in the session
	// and have an intact WS connection
	users    map[string]*User
	maxUsers int
	numUsers int
}

var _ Controller = (*controlBlock)(nil)

// NewControlBlock creates a new Session controlBlock with unique ID.
func NewControlBlock(sessionID string, cache redis.Handler, dispatcher Dispatcher, maxUsers int) *controlBlock {
	scb := &controlBlock{
		id:          sessionID,
		attachments: attachment.NewLocalHandler(sessionID),
		dispatcher:  dispatcher,
		broadcast:   make(chan *types.Message),
		echo:        make(chan *types.Message),
		cache:       cache,
		dbUpdate:    make(chan []redis.Stroke),
		signalClose: make(chan struct{}),
		usersReady:  make(map[string]*User),
		users:       make(map[string]*User),
		maxUsers:    maxUsers,
	}

	// start goroutines for broadcasting and saving changes to board
	go scb.broadcastLoop()
	go scb.dbUpdateLoop()

	return scb
}

// Close sends a close signal
func (scb *controlBlock) Close() {
	scb.signalClose <- struct{}{}
}

func (scb *controlBlock) ID() string {
	return scb.id
}

func (scb *controlBlock) Attachments() attachment.Handler {
	return scb.attachments
}

func (scb *controlBlock) NumUsers() int {
	return scb.numUsers
}

// broadcastLoop Broadcasts board updates to all clients
func (scb *controlBlock) broadcastLoop() {
	for {
		select {
		case data := <-scb.broadcast:
			scb.muUsr.RLock()
			for userID, user := range scb.users { // Send to all connected clients
				// except the origin, i.e. the initiator of message
				if userID != data.Sender {
					if err := user.Conn.WriteJSON(data); err != nil {
						log.Printf("%s :: cannot broadcast to %s: %v",
							scb.id, user.Conn.RemoteAddr(), err)
					}
				}
			}
			scb.muUsr.RUnlock()
		case data := <-scb.echo:
			// echo message back to origin
			scb.muUsr.RLock()
			if err := scb.users[data.Sender].Conn.WriteJSON(data); err != nil {
				log.Printf("error in broadcastLoop: %v", err)
			}
			scb.muUsr.RUnlock()
		case <-scb.signalClose:
			return
		}
	}
}

// dbUpdateLoop updates database according to given Stroke values
func (scb *controlBlock) dbUpdateLoop() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		select {
		case strokes := <-scb.dbUpdate:
			if err := scb.cache.UpdateStrokes(ctx, scb.id, strokes...); err != nil {
				log.Printf("error in dbUpdateLoop: %v", err)
			}
		case <-scb.signalClose:
			_ = scb.cache.ClearSession(ctx, scb.id)
			return
		}
	}
}
