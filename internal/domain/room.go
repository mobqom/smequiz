package domain

import (
	"net/http"
	"sync"

	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

type Room struct {
	id      string
	players map[string]*Player
	screen  dto.Screen
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
func (room *Room) SendMsg(r *http.Request, msg *dto.Msg) {
	// Блокируем чтение карты на время итерации
	room.mu.RLock()
	// Копируем игроков в слайс, чтобы не держать блокировку во время отправки
	players := make([]*Player, 0, len(room.players))
	for _, p := range room.players {
		players = append(players, p)
	}
	room.mu.RUnlock()

	// Отправляем сообщения без блокировки
	for _, p := range players {
		p.SendMsg(r, msg)
	}
}

func (room *Room) SetScreen(r *http.Request, screen dto.Screen) {
	room.SendMsg(
		r,
		&dto.Msg{
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

// GetPlayer возвращает игрока по ID
func (room *Room) GetPlayer(id string) (*Player, bool) {
	room.mu.RLock()
	defer room.mu.RUnlock()
	player, exists := room.players[id]
	return player, exists
}

// Remove метод GetPlayers, так как он небезопасен
