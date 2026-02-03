package domain

type Room interface {
	Join(player Player)
	Leave(player Player)
	PlayersCount() int
	GetPlayers() map[string]Player
}
type room struct {
	id      string
	players map[string]Player
}

func NewRoom(id string) Room {
	return &room{
		id:      id,
		players: make(map[string]Player),
	}
}

func (room *room) Join(player Player) {
	room.players[player.GetId()] = player
}

func (room *room) Leave(player Player) {
	delete(room.players, player.GetId())
}
func (room *room) PlayersCount() int {
	return len(room.players)
}
func (room *room) GetPlayers() map[string]Player {
	return room.players
}
