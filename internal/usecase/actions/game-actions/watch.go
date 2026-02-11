package game_actions

import (
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func Watch(r *http.Request, reqMsg *dto.Msg, p *domain.Player, m *domain.RoomManager) {
	switch reqMsg.Action {
	case dto.START_GAME:
		room, err := m.GetRoom(p.GetRoomId())
		if err != nil {
			log.Printf("room does not exist %s\n", err)
		}
		room.SetScreen(r, dto.TIMER_SCREEN)
		for v := range utils.Timer(3) {
			p := &dto.TimerPayload{Value: v, Done: false}
			if v == 0 {
				p.Done = true
			}
			room.SendMsg(r, &dto.Msg{Action: dto.TIMER_TIME, Payload: p})
		}
		plist := room.GetPlayersSnapshot()
		questions := domain.InitQuestion()
		for ki, pi := range plist {
			for kj, pj := range plist {
				if ki != kj {
					randQuestion := questions[utils.RandRangeInt(0, len(questions))]
					stage := &domain.Stage{
						Question: &randQuestion,
						Players:  map[string]*domain.Player{pi.GetId(): pi, pj.GetId(): pj},
					}
					room.AddStage(stage)
				}
			}
		}

	}
}
