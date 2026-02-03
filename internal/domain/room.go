package domain

import "github.com/ibezgin/mobqom-smequiz/internal/dto"

type Room interface {
	Join(player Player)
	Leave(player Player)
	PlayersCount() int
	GetPlayers() map[string]Player
	SendMsg(msg *dto.Msg)
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
func (room *room) SendMsg(msg *dto.Msg) {
	for _, c := range room.players {
		c.SendMsg(msg)
	}
}
