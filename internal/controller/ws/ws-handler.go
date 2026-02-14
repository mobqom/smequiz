package router

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/usecase"
)

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
	usecase.InitPlayer(r, m, conn)
}
