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
		pCoupleIds := make([]string, 0)
		qList := domain.InitQuestion()
		for k1, p1 := range plist {
			for k2, p2 := range plist {
				if k1 == k2 {
					continue
				}
				cOfP := countOfPlayers(pCoupleIds, k1)
				if cOfP >= 2 {
					continue
				}
				q := qList[utils.RandRangeInt(0, len(qList)-1)]
				pCoupleIds = append(pCoupleIds, k1)
				room.AddStage(&domain.Stage{Players: map[string]*domain.Player{k1: p1, k2: p2}, Question: &q})
			}
		}
		stages := room.GetStage()
		for _, stage := range stages {
			for _, p := range stage.Players {
				p.SendMsg(r, &dto.Msg{Action: dto.SET_SCREEN, Payload: dto.QUESTION_SCREEN})
				p.SendMsg(r, &dto.Msg{Action: dto.SET_QUESTION, Payload: stage.Question})
			}
		}
	}
}

func countOfPlayers(pCoupleIds []string, pId string) int {
	result := 0
	for _, v := range pCoupleIds {
		if v == pId {
			result++
		}
	}
	return result
}
