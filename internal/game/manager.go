package game

type GameManager struct {
	rooms map[string]*Room
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
	gm.rooms[roomID] = newRoom
	return newRoom
}

func (gm *GameManager) GetRoom(roomID string) (*Room, bool) {
	room, exists := gm.rooms[roomID]
	return room, exists
}
