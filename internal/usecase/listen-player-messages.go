package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func ListenMessageLoop(conn *websocket.Conn, m domain.RoomManager) {
	for {
		reqMsg := new(dto.ReqMsg)
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading from websocket:", err)
		}
		err = json.Unmarshal(msgBytes, reqMsg)

		if err != nil {
			fmt.Println("Error unmarshalling from websocket:", err)
		}
		switch reqMsg.Action {
		case dto.CREATE_ROOM:
			roomId := utils.GenerateId("room")
			m.CreateRoom(roomId)

		}
	}
}
