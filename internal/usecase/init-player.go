package usecase

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/usecase/actions"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func InitPlayer(r *http.Request, m *domain.RoomManager, conn *websocket.Conn) {
	playerId := utils.GenerateId("player")
	p := domain.NewPlayer(conn, playerId)
	readPlayerMessages(r, p, m)
}

func readPlayerMessages(r *http.Request, p *domain.Player, m *domain.RoomManager) {
	defer func() {
		roomActions.DeleteEmptyRoom(p, m)
	}()
	for {
		var msg dto.Msg

		err := wsjson.Read(r.Context(), p.Conn(), &msg)
		if err != nil {
			// Нормальное закрытие
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
				websocket.CloseStatus(err) == websocket.StatusGoingAway {
				log.Printf("Клиент отключился: %v", err)
				return
			}
			log.Printf("Ошибка чтения: %v", err)
			return
		}

		actions.StartActionWatchers(r, m, p, msg)
	}
}
