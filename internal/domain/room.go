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
	stages  []*Stage
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
	room.players[player.Id()] = player
}

func (room *Room) Leave(player *Player) {
	room.mu.Lock()
	defer room.mu.Unlock()
	delete(room.players, player.Id())
}
func (room *Room) PlayersCount() int {
	return len(room.players)
}
func (room *Room) SendMsg(ctx context.Context, msg dto.Msg) {
	pls := room.PlayersSnapshot()
	// Отправляем сообщения без блокировки
	for _, p := range pls {
		p.SendMsg(ctx, msg)
	}
}

func (room *Room) SetScreen(ctx context.Context, screen dto.Screen) {
	ps := room.PlayersSnapshot()
	for _, p := range ps {
		go func() {
			p.SetScreen(ctx, screen)
		}()
	}
}

// PlayersSnapshot возвращает копию карты игроков
func (room *Room) PlayersSnapshot() map[string]*Player {
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
	room.stages = append(room.stages, stage)
}

func (room *Room) Stages() []*Stage {
	room.mu.RLock()
	defer room.mu.RUnlock()
	return room.stages
}

func (room *Room) GetStageById(id string) *Stage {
	for _, stage := range room.stages {
		if stage.Id == id {
			return stage
		}
	}
	return nil
}
