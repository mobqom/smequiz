package game_actions

import (
	"context"
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

const (
	MaxPlayersInStage int = 2
)

func Watch(r *http.Request, reqMsg dto.Msg, p *domain.Player, m *domain.RoomManager) {
	switch reqMsg.Action {
	case dto.START_GAME:
		room, err := m.GetRoom(p.GetRoomId())
		if err != nil {
			log.Printf("room does not exist %s\n", err)
		}
		room.SetScreen(r.Context(), dto.TIMER_SCREEN)
		for v := range utils.Timer(3) {
			p := &dto.TimerPayload{Value: v, Done: false}
			if v == 0 {
				p.Done = true
			}
			room.SendMsg(r.Context(), dto.Msg{Action: dto.TIMER_TIME, Payload: p})
		}
		plist := room.GetPlayersSnapshot()
		qList := domain.InitQuestion()

		stages := InitRoomStages(plist, qList, room)
		ctx := context.Background()
		for _, p := range plist {
			go func() {
				stg := findStageWithoutAnswer(stages, p)
				if stg == nil {
					return
				}
				payload := dto.QuestionPayload{Question: stg.Question.Text, StageId: stg.Id}
				p.SendMsg(ctx, dto.Msg{Action: dto.SET_SCREEN, Payload: dto.QUESTION_SCREEN})
				p.SendMsg(ctx, dto.Msg{Action: dto.SET_QUESTION, Payload: payload})
			}()
		}
	}
}

func getPlayersCount(pCoupleIds []string, pId string) int {
	result := 0
	for _, v := range pCoupleIds {
		if v == pId {
			result++
		}
	}
	return result
}

func InitRoomStages(playersList map[string]*domain.Player, qList []domain.Question, room *domain.Room) []*domain.Stage {
	pCoupleIds := make([]string, 0, len(playersList))

	for k1, p1 := range playersList {
		for k2, p2 := range playersList {
			if k1 == k2 {
				continue
			}
			countOfPlayers := getPlayersCount(pCoupleIds, k1)
			if countOfPlayers >= MaxPlayersInStage {
				continue
			}
			q := qList[utils.RandRangeInt(0, len(qList)-1)]
			pCoupleIds = append(pCoupleIds, k1, k2)
			room.AddStage(&domain.Stage{
				Players:  map[string]*domain.Player{k1: p1, k2: p2},
				Question: &q,
				Id:       utils.GenerateId("stage")})
		}
	}

	return room.GetStage()
}

func findStageWithoutAnswer(stage []*domain.Stage, p *domain.Player) *domain.Stage {
	for _, s := range stage {
		_, existP := s.Players[p.GetId()]
		_, existAns := s.Answer[p.GetId()]
		if existP && existAns {
			continue
		}
		return s
	}
	return nil
}
