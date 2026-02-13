package router

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/usecase"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

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

		usecase.StartActionWatchers(r, m, p, msg)
	}
}

func HandleWebSocket(m *domain.RoomManager, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		// Разрешаем любые origins (для разработки)
		InsecureSkipVerify: true,
		// Включаем сжатие
		CompressionMode: websocket.CompressionContextTakeover,
	})
	if err != nil {
		log.Printf("Ошибка Accept: %v", err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "соединение закрыто")
	log.Printf("Клиент подключен: %s", r.RemoteAddr)

	playerId := utils.GenerateId("player")
	p := domain.NewPlayer(conn, playerId)
	readPlayerMessages(r, p, m)
}
