package game_actions

import (
	"log"
	"net/http"
	"time"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

func Watch(r *http.Request, reqMsg *dto.Msg, p *domain.Player, m *domain.RoomManager) {
	switch reqMsg.Action {
	case dto.START_GAME:
		room, err := m.GetRoom(p.GetRoomId())
		if err != nil {
			log.Printf("room does not exist %s\n", err)
		}
		room.SetScreen(r, dto.TIMER_SCREEN)
		go func() {
			for i := 3; i > 0; i-- {
				msg := &dto.Msg{Action: dto.TIMER_TIME, Payload: i}
				room.SendMsg(r, msg)
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
