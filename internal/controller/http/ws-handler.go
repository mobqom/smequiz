package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/usecase"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
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
		return
	}

	playerId := utils.GenerateId("player")
	p := domain.NewPlayer(conn, playerId)
	fmt.Println("Client join the server", p.GetConn().LocalAddr().String())
	go usecase.ReadPlayerMessages(p, m)
}
