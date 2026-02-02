package game

import "sync"

type GameManager struct {
	rooms map[string]*Room
	mu    sync.Mutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		rooms: make(map[string]*Room),
	}
}

func (gm *GameManager) GetOrCreateRoom(roomID string) *Room {
	if room, exists := gm.GetRoom(roomID); exists {
		return room
	}
	newRoom := NewRoom(roomID)
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.rooms[roomID] = newRoom
	return newRoom
}

func (gm *GameManager) GetRoom(roomID string) (*Room, bool) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	room, exists := gm.rooms[roomID]
	return room, exists
}

func (gm *GameManager) DeleteRoom(roomID string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	delete(gm.rooms, roomID)
}
