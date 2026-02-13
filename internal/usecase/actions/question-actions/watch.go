package questionActions

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	game_actions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/game-actions"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func Watch(r *http.Request, reqMsg dto.Msg, m *domain.RoomManager, p *domain.Player) {
	switch reqMsg.Action {
	case dto.ANSWER_QUESTION:
		var mu sync.Mutex
		validate := validator.New()
		var data dto.AnswerPayload
		err := utils.PreparePayloadToStruct(reqMsg.Payload, &data)
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
			p.SendMsg(r.Context(), dto.Msg{Action: dto.SET_SCREEN, Payload: dto.QUESTION_RESULT_SCREEN})
			return
		}
		payload := dto.QuestionPayload{Question: swa.Question.Text, StageId: swa.Id}

		p.SendMsg(r.Context(), dto.Msg{Action: dto.SET_QUESTION, Payload: payload})
	}
}
