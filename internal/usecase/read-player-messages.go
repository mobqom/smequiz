package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	roomActions "github.com/ibezgin/mobqom-smequiz/internal/usecase/actions"
)

func ReadPlayerMessages(p domain.Player, m domain.RoomManager) {
	defer func() {
		fmt.Println("Client left the server", p.GetConn().LocalAddr().String())
	}()
	for {
		reqMsg := new(dto.ReqMsg)
		_, msgBytes, err := p.GetConn().ReadMessage()
		if err != nil {
			fmt.Println("Error reading from websocket:", err)
		}
		err = json.Unmarshal(msgBytes, reqMsg)

		if err != nil {
			fmt.Println("Error unmarshalling from websocket:", err)
		}

		roomActions.Watch(m, p, reqMsg)
	}
}
