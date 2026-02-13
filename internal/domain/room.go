package domain

import (
	"context"
	"sync"

	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

type Room struct {
	id      string
	players map[string]*Player
	screen  dto.Screen
	stage   []*Stage
	mu      sync.RWMutex
}

func NewRoom(id string) *Room {
	return &Room{
		id:      id,
		players: make(map[string]*Player),
		screen:  dto.WAITING_SCREEN,
	}
}
func (room *Room) Join(player *Player) {
	room.mu.Lock()
	defer room.mu.Unlock()
	room.players[player.GetId()] = player
}

func (room *Room) Leave(player *Player) {
	room.mu.Lock()
	defer room.mu.Unlock()
	delete(room.players, player.GetId())
}
func (room *Room) PlayersCount() int {
	return len(room.players)
}
func (room *Room) SendMsg(ctx context.Context, msg dto.Msg) {
	pls := room.GetPlayersSnapshot()
	// Отправляем сообщения без блокировки
	for _, p := range pls {
		p.SendMsg(ctx, msg)
	}
}

func (room *Room) SetScreen(ctx context.Context, screen dto.Screen) {
	room.SendMsg(
		ctx,
		dto.Msg{
			Action:  dto.SET_SCREEN,
			Payload: screen,
		})
}

// GetPlayersSnapshot возвращает копию карты игроков
func (room *Room) GetPlayersSnapshot() map[string]*Player {
	room.mu.RLock()
	defer room.mu.RUnlock()

	// Создаем копию карты
	playersCopy := make(map[string]*Player, len(room.players))
	for id, player := range room.players {
		playersCopy[id] = player
	}
	return playersCopy
}
func (room *Room) GetPlayers() map[string]*Player {
	room.mu.Lock()
	defer room.mu.Unlock()
	return room.players
}

func (room *Room) AddStage(stage *Stage) {
	room.mu.Lock()
	defer room.mu.Unlock()
	room.stage = append(room.stage, stage)
}

func (room *Room) GetStage() []*Stage {
	room.mu.RLock()
	defer room.mu.RUnlock()
	return room.stage
}
