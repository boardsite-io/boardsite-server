package session

import (
	"encoding/json"
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/heat1q/boardsite/api/types"
	"github.com/heat1q/boardsite/redis"
)

var receiveHandler = map[string]func(scb *ControlBlock, msg *types.Message) error{
	types.MessageTypeStroke:    sanitizeStrokes,
	types.MessageTypeMouseMove: mouseMove,
}

// GetStrokes fetches all stroke data for specified page.
func GetStrokes(sessionID, pageID string) ([]types.Stroke, error) {
	strokesRaw, err := redis.FetchStrokesRaw(sessionID, pageID)
	if err != nil {
		return nil, errors.New("unable to fetch strokes")
	}

	strokes := make([]types.Stroke, len(strokesRaw))
	for i, s := range strokesRaw {
		if err := json.Unmarshal(s, &strokes[i]); err != nil {
			return nil, err
		}
	}
	return strokes, nil
}

// GetPages returns all pageIDs in order.
func GetPages(sessionID string) ([]string, map[string]*types.PageMeta, error) {
	pageRank, err := redis.GetPages(sessionID)
	if err != nil {
		return nil, nil, errors.New("unable to fetch pages")
	}
	meta, err := redis.GetPagesMeta(sessionID, pageRank...)
	if err != nil {
		return nil, nil, errors.New("unable to fetch pages meta data")
	}
	return pageRank, meta, nil
}

// GetPagesSet returns all pageIDs in a map for fast verification.
func GetPagesSet(sessionID string) map[string]struct{} {
	pageIDs, _ := redis.GetPages(sessionID)
	pageIDSet := make(map[string]struct{})

	for _, pid := range pageIDs {
		pageIDSet[pid] = struct{}{}
	}

	return pageIDSet
}

// IsValidPage checks if a pageID is valid, i.e. the page exists.
func IsValidPage(sessionID string, pageID ...string) bool {
	pages := GetPagesSet(sessionID)
	for _, pid := range pageID {
		if _, ok := pages[pid]; !ok {
			return false
		}
	}
	return true
}

// AddPages adds a page with pageID to the session and broadcasts
// the change to all connected clients.
func AddPages(scb *ControlBlock, pageIDs []string, index []int, meta map[string]*types.PageMeta) error {
	if len(pageIDs) != len(index) {
		return errors.New("cannot find page index")
	}
	if IsValidPage(scb.ID, pageIDs...) {
		return errors.New("some pages already exist")
	}

	defer SyncPages(scb)

	for i, pid := range pageIDs {
		pMeta, ok := meta[pid]
		if !ok {
			return fmt.Errorf("no meta given for page %s", pid)
		}
		if err := redis.AddPage(scb.ID, pid, index[i], pMeta); err != nil {
			return errors.New("cannot add page")
		}
	}

	return nil
}

// DeletePage deletes a page with pageID and broadcasts
// the change to all connected clients.
func DeletePage(scb *ControlBlock, pageID string) error {
	//TODO handle error
	if err := redis.DeletePage(scb.ID, pageID); err != nil {
		return errors.New("cannot delete page")
	}
	return SyncPages(scb)
}

// UpdatePages modifies the page meta data and/or clears the content.
func UpdatePages(scb *ControlBlock, pageIDs []string, meta map[string]*types.PageMeta, clear bool) error {
	var pageIDsUpdate []string

	defer func() {
		if len(pageIDsUpdate) == 0 {
			return
		}
		scb.Broadcast <- &types.Message{
			Type:   types.MessageTypePageUpdate,
			Sender: "", // send to all clients
			Content: types.ContentPageRequest{
				PageID: pageIDsUpdate,
				Clear:  clear,
				Meta:   meta,
			},
		}
	}()

	for _, pid := range pageIDs {
		if !IsValidPage(scb.ID, pid) {
			return fmt.Errorf("page %s does not exits", pid)
		}
		if clear {
			if err := redis.ClearPage(scb.ID, pid); err != nil {
				return fmt.Errorf("cannot clear page %s", pid)
			}
		} else {
			pMeta, ok := meta[pid]
			if !ok {
				return fmt.Errorf("no meta given for page %s", pid)
			}
			// update db
			if err := redis.UpdatePageMeta(scb.ID, pid, pMeta); err != nil {
				return err
			}
		}
		pageIDsUpdate = append(pageIDsUpdate, pid)
	}

	return nil
}

// SyncPages broadcasts the current PageRank to all connected
// clients indicating an update in the pages (or ordering).
func SyncPages(scb *ControlBlock) error {
	pageRank, meta, err := GetPages(scb.ID)
	if err != nil {
		return err
	}

	scb.Broadcast <- &types.Message{
		Type:   types.MessageTypePageSync,
		Sender: "", // send to all clients
		Content: &types.ContentPageSync{
			PageRank: pageRank,
			Meta:     meta,
		},
	}
	return nil
}

// NewUser generate a new user struct based on
// the alias and color attribute
//
// Does some sanitize checks.
func NewUser(scb *ControlBlock, alias, color string) (*types.User, error) {
	if len(alias) > 24 {
		alias = alias[:24]
	}
	//TODO check if html color ?
	if len(color) != 7 {
		return nil, fmt.Errorf("incorrect html color")
	}

	id, err := gonanoid.New(16)
	if err != nil {
		return nil, err
	}
	user := &types.User{
		ID:    id,
		Alias: alias,
		Color: color,
	}
	// set user waiting
	scb.UserReady(user)
	return user, err
}

// Receive is the entry point when a message is received in
// the session via the websocket.
func Receive(scb *ControlBlock, msg *types.Message) error {
	if !scb.IsUserConnected(msg.Sender) {
		return errors.New("invalid sender userId")
	}
	handleMsg, ok := receiveHandler[msg.Type]
	if !ok {
		return fmt.Errorf("message type not recognized: %s", msg.Type)
	}
	return handleMsg(scb, msg)
}

// sanitizeStrokes parses the stroke content of the message.
//
// It further checks if the strokes have a valid pageId and userId.
func sanitizeStrokes(scb *ControlBlock, msg *types.Message) error {
	var strokes []*types.Stroke
	if err := msg.UnmarshalContent(&strokes); err != nil {
		return err
	}

	validStrokes := make([]*types.Stroke, 0, len(strokes))
	pageIDs := GetPagesSet(scb.ID)

	for _, stroke := range strokes {
		if _, ok := pageIDs[stroke.GetPageID()]; ok { // valid pageID
			if stroke.GetUserID() == msg.Sender { // valid userID
				validStrokes = append(validStrokes, stroke)
			}
		}
	}
	if len(validStrokes) > 0 {
		updateStrokes(scb, msg.Sender, validStrokes)
		return nil
	}
	return errors.New("strokes not validated")
}

// updateStrokes updates the strokes in the session with sessionID.
//
// userID indicates the initiator of the message, which is
// to be excluded in the broadcast. The strokes are scheduled for an
// update to Redis.
func updateStrokes(scb *ControlBlock, userID string, strokes []*types.Stroke) {
	// broadcast changes
	scb.Broadcast <- &types.Message{
		Type:    types.MessageTypeStroke,
		Sender:  userID,
		Content: strokes,
	}

	// save to database
	scb.DBUpdate <- strokes
}

// mouseMove broadcast mouse move events.
func mouseMove(scb *ControlBlock, msg *types.Message) error {
	var mouseUpdate types.ContentMouseMove
	if err := msg.UnmarshalContent(&mouseUpdate); err != nil {
		return err
	}
	scb.Broadcast <- &types.Message{
		Type:    types.MessageTypeMouseMove,
		Sender:  msg.Sender,
		Content: mouseUpdate,
	}
	return nil
}
