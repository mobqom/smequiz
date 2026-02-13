package questionActions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	game_actions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/game-actions"
)

func Watch(r *http.Request, reqMsg dto.Msg, m *domain.RoomManager, p *domain.Player) {
	switch reqMsg.Action {
	case dto.ANSWER_QUESTION:
		var mu sync.Mutex
		validate := validator.New()
		var data dto.AnswerPayload
		err := preparePayloadToStruct(reqMsg.Payload, &data)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
		err = validate.Struct(data)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		room, _ := m.RoomById(p.RoomId())
		s := room.GetStageById(data.StageId)
		if s == nil {
			fmt.Printf("err: stage not exist\n")
			return
		}
		mu.Lock()
		s.Answer[p.Id()] = data.Answer
		mu.Unlock()
		slist := room.Stages()
		swa := game_actions.FindStageWithoutAnswer(slist, p)
		if swa == nil {
			p.SendMsg(r.Context(), dto.Msg{Action: dto.SET_SCREEN, Payload: dto.WAITING_SCREEN})
			return
		}
		payload := dto.QuestionPayload{Question: swa.Question.Text, StageId: swa.Id}

		p.SendMsg(r.Context(), dto.Msg{Action: dto.SET_QUESTION, Payload: payload})
	}
}

func preparePayloadToStruct[T any](payload interface{}, data *T) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling payload: %v", err)
		return err
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("Error unmarshaling to AnswerPayload: %v", err)
		return err
	}
	return nil
}
