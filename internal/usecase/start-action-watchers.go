package usecase

import (
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	questionActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/question-actions"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
)

func StartActionWatchers(m domain.RoomManager, p domain.Player, reqMsg *dto.Msg) {
	go roomActions.Watch(m, p, reqMsg)
	go questionActions.Watch()
}
