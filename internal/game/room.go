package game

type Room struct {
	ID      string
	players map[string]*Player
}

func NewRoom(id string) *Room {
	return &Room{
		ID:      id,
		players: map[string]*Player{},
	}
}

func (r *Room) AddPlayer(p *Player) {
	r.players[p.ID] = p
}

func (r *Room) RemovePlayer(p *Player) {
	delete(r.players, p.ID)
}
