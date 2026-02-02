package game

type gameManager struct {
	rooms map[string]*room
}

func NewGameManager() *gameManager {
	return &gameManager{
		rooms: map[string]*room{},
	}
}
