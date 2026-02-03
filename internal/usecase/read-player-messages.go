package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions"
)

func ReadPlayerMessages(p domain.Player, m domain.RoomManager) {
	defer func() {
		fmt.Println("Client left the server", p.GetConn().LocalAddr().String())
		p.GetConn().Close()
	}()
	for {
		reqMsg := new(dto.ReqMsg)
		_, msgBytes, err := p.GetConn().ReadMessage()
		if err != nil {
			// Normal closure (1000) or going away (1001) are expected
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Println("Error reading from websocket:", err)
			}
			return
		}
		err = json.Unmarshal(msgBytes, reqMsg)

		if err != nil {
			fmt.Println("Error unmarshalling from websocket:", err)
			return
		}
		roomActions.Watch(m, p, reqMsg)
	}
}
