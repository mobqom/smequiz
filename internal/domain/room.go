package domain

import "github.com/ibezgin/mobqom-smequiz/internal/dto"

type Room struct {
	id      string
	players map[string]*Player
	screen  dto.Screen
}

func NewRoom(id string) *Room {
	return &Room{
		id:      id,
		players: make(map[string]*Player),
		screen:  dto.WAITING_SCREEN,
	}
}
func (room *Room) Join(player *Player) {
	room.players[player.GetId()] = player
}

func (room *Room) Leave(player *Player) {
	delete(room.players, player.GetId())
}
func (room *Room) PlayersCount() int {
	return len(room.players)
}
func (room *Room) SendMsg(msg *dto.Msg) {
	for _, c := range room.players {
		c.SendMsg(msg)
	}
}
func (room *Room) SetScreen(screen dto.Screen) {
	room.SendMsg(&dto.Msg{
		Action:  dto.SET_SCREEN,
		Payload: screen,
	})
}
func (room *Room) GetPlayers() map[string]*Player {
	return room.players
}
