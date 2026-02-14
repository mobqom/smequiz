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
	TimeToEndRound        = 30
)

func Watch(r *http.Request, reqMsg dto.Msg, p *domain.Player, m *domain.RoomManager) {
	switch reqMsg.Action {
	case dto.START_GAME:
		go func() {
			stagesCh := make(chan []*domain.Stage)
			room, err := m.RoomById(p.RoomId())
			if err != nil {
				log.Printf("room does not exist %s\n", err)
			}
			roomPlayers := room.PlayersSnapshot()

			go func() {
				qList := domain.InitQuestion()
				stagesCh <- InitRoomStages(roomPlayers, qList, room)
			}()

			room.SetScreen(r.Context(), dto.TIMER_SCREEN)
			SetTimerTime(r.Context(), room, 3)

			stages := <-stagesCh
			for _, p := range roomPlayers {
				go func() {
					stg := FindStageWithoutAnswer(stages, p)
					if stg == nil {
						return
					}
					payload := dto.QuestionPayload{Question: stg.Question.Text, StageId: stg.Id}
					p.SetScreen(r.Context(), dto.QUESTION_SCREEN)
					p.SendMsg(r.Context(), dto.Msg{Action: dto.SET_QUESTION, Payload: payload})
				}()
			}
			WaitRoundEnd(r, room)
		}()
	}
}

func WaitRoundEnd(r *http.Request, room *domain.Room) {
	cancelTimerCtx, cancel := context.WithCancel(context.Background())

	go func() {
		SetTimerTime(cancelTimerCtx, room, TimeToEndRound)
		room.SetScreen(r.Context(), dto.QUESTION_RESULT_SCREEN)
	}()

	go func() {
		stages := room.Stages()
		isAllAnswersDone := true
		for {
			select {
			case <-cancelTimerCtx.Done():
				room.SetScreen(r.Context(), dto.QUESTION_RESULT_SCREEN)
				return
			default:
				isAllAnswersDone = CheckStageToAllAnswers(stages)
			}
			if isAllAnswersDone {
				cancel()
			}
		}
	}()

}

func CheckStageToAllAnswers(stages []*domain.Stage) bool {
	result := true
	for _, s := range stages {
		ansCount := len(s.Answer)
		if ansCount < MaxPlayersInStage {
			result = false
		}
	}
	return result
}
func SetTimerTime(ctx context.Context, room *domain.Room, seconds int) {
	for v := range utils.Timer(seconds) {
		select {
		case <-ctx.Done():
			return
		default:
			p := &dto.TimerPayload{Value: v, Done: false}
			if v == 0 {
				p.Done = true
			}
			room.SendMsg(ctx, dto.Msg{Action: dto.TIMER_TIME, Payload: p})
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

func InitRoomStages(roomPlayers map[string]*domain.Player, qList []domain.Question, room *domain.Room) []*domain.Stage {
	pCoupleIds := make([]string, 0, len(roomPlayers)*MaxPlayersInStage)
	for k1, p1 := range roomPlayers {
		for k2, p2 := range roomPlayers {
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
