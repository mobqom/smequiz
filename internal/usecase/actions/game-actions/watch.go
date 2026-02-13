package game_actions

import (
	"context"
	"fmt"
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
		stagesCh := make(chan []*domain.Stage)
		room, err := m.RoomById(p.RoomId())
		if err != nil {
			log.Printf("room does not exist %s\n", err)
		}
		room.SetScreen(r.Context(), dto.TIMER_SCREEN)
		plist := room.PlayersSnapshot()

		go func() {
			qList := domain.InitQuestion()

			stagesCh <- InitRoomStages(plist, qList, room)
		}()
		for v := range utils.Timer(3) {
			p := &dto.TimerPayload{Value: v, Done: false}
			if v == 0 {
				p.Done = true
			}
			room.SendMsg(r.Context(), dto.Msg{Action: dto.TIMER_TIME, Payload: p})
		}

		ctx := context.Background()
		stages := <-stagesCh
		for _, p := range plist {
			go func() {
				stg := FindStageWithoutAnswer(stages, p)
				if stg == nil {
					return
				}
				payload := dto.QuestionPayload{Question: stg.Question.Text, StageId: stg.Id}
				p.SetScreen(ctx, dto.QUESTION_SCREEN)
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
	fmt.Println(pCoupleIds)
	for k1, p1 := range playersList {
		for k2, p2 := range playersList {
			if k1 == k2 {
				continue
			}
			countOfPlayers := getPlayersCount(pCoupleIds, k1)
			if countOfPlayers >= MaxPlayersInStage {
				break
			}
			q := qList[utils.RandRangeInt(0, len(qList)-1)]
			pCoupleIds = append(pCoupleIds, k1, k2)
			room.AddStage(&domain.Stage{
				Players:  map[string]*domain.Player{k1: p1, k2: p2},
				Question: &q,
				Id:       utils.GenerateId("stage")})
		}
	}

	return room.Stages()
}

// FindStageWithoutAnswer Если в стейдже есть плеер, но нет ответа от него возвращает стейдж
func FindStageWithoutAnswer(stage []*domain.Stage, p *domain.Player) *domain.Stage {
	for _, s := range stage {
		_, existP := s.Players[p.Id()]
		_, existAns := s.Answer[p.Id()]
		if existP && !existAns {
			return s
		}
	}
	return nil
}
