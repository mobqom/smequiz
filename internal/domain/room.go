package domain

type Room interface {
	Join(player Player)
	Leave(player Player)
}
type room struct {
	id      string
	clients map[string]*Player
}

func NewRoom(id string) Room {
	return &room{
		id:      id,
		clients: make(map[string]*Player),
	}
}

func (room *room) Join(player Player) {
	room.clients[player.GetId()] = &player
}

func (room *room) Leave(player Player) {}
