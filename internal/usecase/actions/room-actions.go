package roomActions

import (
	"fmt"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func Watch(m domain.RoomManager, p domain.Player, reqMsg *dto.ReqMsg) {
	switch reqMsg.Action {
	case dto.CREATE_ROOM:
		roomId := utils.GenerateId("room")
		room, err := m.CreateRoom(roomId)
		if err != nil {
			fmt.Printf("error with create room; %s", err)
			return
		}
		p.SetRoomId(roomId)
		room.Join(p)
		fmt.Printf("room %s has been created")

	case dto.JOIN_ROOM:
		roomId := reqMsg.Payload.(string)
		room, err := m.GetRoom(roomId)
		if err != nil {
			fmt.Printf("room does not exist %s\n", err)
			return
		}
		p.SetRoomId(roomId)
		room.Join(p)
	case dto.LEAVE_ROOM:
		roomId := p.GetRoomId()
		room, err := m.GetRoom(roomId)
		if err != nil {
			fmt.Println("room does not exist")
		}
		room.Leave(p)
		p.SetRoomId("")
	}
}
