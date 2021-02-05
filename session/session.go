package session

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	gws "github.com/gorilla/websocket"

	"github.com/heat1q/boardsite/api/types"
	"github.com/heat1q/boardsite/database"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	// ActiveSession maps the session is to the SessionControl struct
	ActiveSession = make(map[string]*ControlBlock)
)

// GenerateID generates a unique sessionID.
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, 6)
	// find available id
	for {
		for i := range id {
			id[i] = letters[rand.Intn(len(letters))]
		}

		if ActiveSession[string(id)] == nil {
			break
		}
	}

	return string(id)
}

// Create creates and initializes a new SessionControl struct
func Create() string {
	scb := NewControlBlock(GenerateID())

	// assign to SessionControl struct
	ActiveSession[scb.ID] = scb

	// start goroutines for broadcasting and saving changes to board
	go scb.broadcast()
	go scb.updateDatabase()

	log.Printf("Create Session with ID: %s\n", scb.ID)

	return scb.ID
}

// IsValid checks if session with sessionID exists.
func IsValid(sessionID string) bool {
	return ActiveSession[sessionID] != nil
}

// GetPages returns all pageIDs in order.
func GetPages(sessionID string) []string {
	if db, err := database.NewRedisConn(sessionID); err == nil {
		defer db.Close()
		return db.GetPages()
	}
	return []string{}
}

// GetPagesSet returns all pageIDs in a map for fast verification.
func GetPagesSet(sessionID string) map[string]struct{} {
	pageIDs := GetPages(sessionID)
	pageIDSet := make(map[string]struct{})

	for _, pid := range pageIDs {
		pageIDSet[pid] = struct{}{}
	}

	return pageIDSet
}

// IsValidPage checks if a pageID is valid, i.e. the page exists.
func IsValidPage(sessionID, pageID string) bool {
	_, ok := GetPagesSet(sessionID)[pageID]
	return ok
}

// AddPage adds a page with pageID to the session and broadcasts
// the change to all connected clients.
func AddPage(sessionID, pageID string, index int) {
	if db, err := database.NewRedisConn(sessionID); err == nil {
		defer db.Close()
		db.AddPage(pageID, index)
		UpdatePages(
			sessionID,
			db.GetPages(),
			[]string{},
		)
	}
}

// DeletePage deletes a page with pageID and broadcasts
// the change to all connected clients.
func DeletePage(sessionID, pageID string) {
	if db, err := database.NewRedisConn(sessionID); err == nil {
		defer db.Close()
		db.DeletePage(pageID)
		UpdatePages(
			sessionID,
			db.GetPages(),
			[]string{},
		)
	}
}

// ClearPage clears all strokes on page with pageID and broadcasts
// the change to all connected clients.
func ClearPage(sessionID, pageID string) {
	if db, err := database.NewRedisConn(sessionID); err == nil {
		defer db.Close()
		db.ClearPage(pageID)
		UpdatePages(
			sessionID,
			[]string{},
			[]string{pageID},
		)
	}
}

// Close closes a session.
func Close(sessionID string) {
	ActiveSession[sessionID].close()
	delete(ActiveSession, sessionID)
}

// AddClient adds the client to the session
func AddClient(sessionID string, conn *gws.Conn) {
	ActiveSession[sessionID].Mu.Lock()

	// add current remote connections to clients
	ActiveSession[sessionID].NumClients++
	ActiveSession[sessionID].Clients[conn.RemoteAddr().String()] = conn

	ActiveSession[sessionID].Mu.Unlock()
}

// RemoveClient removes the client from the session
func RemoveClient(sessionID, remoteAddr string) {
	ActiveSession[sessionID].Mu.Lock()

	// remove current remote connection from clients
	ActiveSession[sessionID].NumClients--
	delete(ActiveSession[sessionID].Clients, remoteAddr)

	ActiveSession[sessionID].Mu.Unlock()

	// if session is empty after client disconnect
	// the session needs to be set to inactive
	if ActiveSession[sessionID].NumClients == 0 {
		Close(sessionID)
	}
}

// SendAllToClient schedules a request to send all data of a session to client.
func SendAllToClient(sessionID, remoteAddr string) {
	ActiveSession[sessionID].DBFetch <- remoteAddr
}

// UpdatePages broadcasts the current PageRank to all connected
// clients indicating an update in the pages (or ordering).
func UpdatePages(sessionID string, pageIDsToUpdate, pageIDsToClear []string) {
	stroke := types.Stroke{
		Type:      -1, // non-zero, since it's no deletion
		PageRank:  pageIDsToUpdate,
		PageClear: pageIDsToClear,
	}

	UpdateStrokes(
		sessionID,
		"", // send to all clients
		[]*types.Stroke{&stroke},
		[]*types.Stroke{},
	)
}

// UpdateStrokes updates the data in the session with sessionID.
//
// RemoteAddr indicates the initiator of the message, which is
// to be excluded in the broadcast.
// Strokes in the first slice are broadcasted to all connected
// clients. Stroke in the second slice (those with type >= 0)
// are updated in Redis.
func UpdateStrokes(sessionID, remoteAddr string, strokes, strokesDB []*types.Stroke) {
	// ignore error
	// since it is unlikely that marshalling fails
	strokesEncoded, _ := json.Marshal(&strokes)
	// broadcast changes
	ActiveSession[sessionID].Broadcast <- &BroadcastData{
		Origin:  remoteAddr,
		Content: strokesEncoded,
	}

	// save to database
	if len(strokesDB) > 0 {
		ActiveSession[sessionID].DBUpdate <- strokesDB
	}
}

// SanitizeAndRelay sanitizes websocket input data and returns an
// error if data is corrupted.
func SanitizeAndRelay(sessionID, remoteAddr string, data []byte) error {
	var strokes = []types.Stroke{}
	if err := json.Unmarshal(data, &strokes); err != nil {
		return err
	}

	strokesSanitized, strokesDB := SanitizeStrokes(sessionID, strokes)
	// update the session data
	UpdateStrokes(sessionID, remoteAddr, strokesSanitized, strokesDB)
	return nil
}

// SanitizeStrokes sanitizes a slice of strokes.
//
// It divides the input slice into two slices of pointer
// to strokes to prevent hardcopies. The first contains all
// sanitizes slices. The second contains only the stroke that
// need to be stored in the DB (i.e. type >= 0).
func SanitizeStrokes(sessionID string, strokes []types.Stroke) ([]*types.Stroke, []*types.Stroke) {
	strokesSanitized := make([]*types.Stroke, 0, len(strokes))
	strokesDB := make([]*types.Stroke, 0, len(strokes))
	pageIDs := GetPagesSet(sessionID)

	for i := range strokes {
		if _, ok := pageIDs[strokes[i].GetPageID()]; ok {
			strokesSanitized = append(strokesSanitized, &strokes[i])
			if strokes[i].Type >= 0 {
				strokesDB = append(strokesDB, &strokes[i])
			}
		}
	}

	return strokesSanitized, strokesDB
}
