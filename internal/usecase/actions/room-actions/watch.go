package roomActions

import (
	"fmt"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func Watch(m *domain.RoomManager, p *domain.Player, reqMsg *dto.Msg) {
	switch reqMsg.Action {
	case dto.CREATE_ROOM:
		if playerRoomId := p.GetRoomId(); playerRoomId != "" {
			fmt.Printf("player already has room\n")
			return
		}
		roomId := utils.GenerateId("room")
		room, err := m.CreateRoom(roomId)
		if err != nil {
			fmt.Printf("error with create room; %s\n", err)
			return
		}
		p.SetRoomId(roomId)
		room.Join(p)
		fmt.Printf("%s: created room %s\n", p.GetId(), roomId)
		go sendPlayersList(room)
		go sendCurrentRoom(p)
	case dto.JOIN_ROOM:
		roomId := reqMsg.Payload.(string)
		room, err := m.GetRoom(roomId)
		if err != nil {
			fmt.Printf("%s: room does not exist %s\n", roomId, err)
			return
		}
		p.SetRoomId(roomId)
		room.Join(p)
		go sendPlayersList(room)
		go sendCurrentRoom(p)
		fmt.Printf("player %s has bean joined to to room %s\n", p.GetId(), roomId)
	case dto.LEAVE_ROOM:
		roomId := p.GetRoomId()
		room, err := m.GetRoom(roomId)
		if err != nil {
			fmt.Printf("room does not exist %s\n", err)
			return
		}
		room.Leave(p)
		go sendPlayersList(room)
		DeleteEmptyRoom(p, m)
		go sendCurrentRoom(p)
		fmt.Printf("player %s left the room %s\n", p.GetId(), p.GetRoomId())
		p.SetRoomId("")

	default:
	}
}

func DeleteEmptyRoom(p *domain.Player, m *domain.RoomManager) {
	roomId := p.GetRoomId()
	room, err := m.GetRoom(roomId)
	if err != nil {
		return
	}
	room.Leave(p)
	playersCount := room.PlayersCount()
	if playersCount == 0 {
		m.DeleteRoom(roomId)
		fmt.Printf("empty room has been deleted %s\n", roomId)
	}
}

func sendPlayersList(room *domain.Room) {
	var list []string
	players := room.GetPlayersSnapshot()
	for _, c := range players {
		list = append(list, c.GetId())

	}
	room.SendMsg(&dto.Msg{Action: dto.PLAYERS_LIST, Payload: list})
}

func sendCurrentRoom(p *domain.Player) {
	p.SendMsg(
		&dto.Msg{
			Action:  dto.CURRENT_ROOM,
			Payload: p.GetRoomId(),
		},
	)
}
