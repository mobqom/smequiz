package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/usecase"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions/room-actions"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func readPlayerMessages(p domain.Player, m domain.RoomManager) {
	defer func() {
		roomActions.DeleteEmptyRoom(p, m)
		fmt.Println("Client left the server", p.GetId())
	}()
	for {
		_, msgBytes, err := p.GetConn().ReadMessage()
		if err != nil {
			// Normal closure (1000), going away (1001), and no status (1005) are expected
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
				fmt.Println("Error reading from websocket:", err)
			}
			return
		}
		reqMsg := new(dto.ReqMsg)
		err = json.Unmarshal(msgBytes, reqMsg)

		if err != nil {
			fmt.Println("Error unmarshalling from websocket:", err)
			return
		}
		usecase.ActionsWatchersHub(m, p, reqMsg)
	}
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
	go readPlayerMessages(p, m)
}
