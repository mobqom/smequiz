package usecase

import (
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	questionActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/question-actions"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
)

func ActionsWatchersHub(m domain.RoomManager, p domain.Player, reqMsg *dto.ReqMsg) {
	go roomActions.Watch(m, p, reqMsg)
	go questionActions.Watch()
}
