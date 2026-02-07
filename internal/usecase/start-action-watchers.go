package usecase

import (
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	game_actions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/game-actions"
	playeraction "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/player-action"
	questionActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/question-actions"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
)

func StartActionWatchers(m domain.RoomManager, p domain.Player, reqMsg *dto.Msg) {
	go roomActions.Watch(m, p, reqMsg)
	go questionActions.Watch()
	go playeraction.Watch(reqMsg, p)
	go game_actions.Watch(reqMsg, p, m)
}
