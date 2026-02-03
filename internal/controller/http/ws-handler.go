package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/usecase"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWs(m domain.RoomManager, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
	}

	fmt.Println("Connected to WebSocket")
	go usecase.ListenMessageLoop(conn, m)
}
