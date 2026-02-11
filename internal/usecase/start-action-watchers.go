package usecase

import (
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	game_actions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/game-actions"
	playeraction "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/player-action"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
)

func StartActionWatchers(r *http.Request, m *domain.RoomManager, p *domain.Player, reqMsg *dto.Msg) {
	roomActions.Watch(r, reqMsg, m, p)
	playeraction.Watch(r, reqMsg, p)
	game_actions.Watch(r, reqMsg, p, m)
}
