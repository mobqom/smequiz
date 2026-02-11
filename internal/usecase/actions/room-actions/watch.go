package roomActions

import (
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/utils"
)

func Watch(r *http.Request, reqMsg *dto.Msg, m *domain.RoomManager, p *domain.Player) {
	switch reqMsg.Action {
	case dto.CREATE_ROOM:
		if playerRoomId := p.GetRoomId(); playerRoomId != "" {
			log.Printf("player already has room\n")
			return
		}
		roomId := utils.GenerateId("room")
		room, err := m.CreateRoom(roomId)
		if err != nil {
			log.Printf("error with create room; %s\n", err)
			return
		}
		p.SetRoomId(roomId)
		room.Join(p)
		log.Printf("%s: created room %s\n", p.GetId(), roomId)
		go sendPlayersList(r, room)
		go sendCurrentRoom(r, p)
	case dto.JOIN_ROOM:
		roomId := reqMsg.Payload.(string)
		room, err := m.GetRoom(roomId)
		if err != nil {
			log.Printf("%s: room does not exist %s\n", roomId, err)
			return
		}
		p.SetRoomId(roomId)
		room.Join(p)
		go sendPlayersList(r, room)
		go sendCurrentRoom(r, p)
		log.Printf("player %s has bean joined to to room %s\n", p.GetId(), roomId)
	case dto.LEAVE_ROOM:
		roomId := p.GetRoomId()
		room, err := m.GetRoom(roomId)
		if err != nil {
			log.Printf("room does not exist %s\n", err)
			return
		}
		room.Leave(p)
		go sendPlayersList(r, room)
		DeleteEmptyRoom(p, m)
		go sendCurrentRoom(r, p)
		log.Printf("player %s left the room %s\n", p.GetId(), p.GetRoomId())
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
		log.Printf("empty room has been deleted %s\n", roomId)
	}
}

func sendPlayersList(r *http.Request, room *domain.Room) {
	var list []string
	players := room.GetPlayersSnapshot()
	for _, c := range players {
		list = append(list, c.GetId())

	}
	room.SendMsg(r, &dto.Msg{Action: dto.PLAYERS_LIST, Payload: list})
}

func sendCurrentRoom(r *http.Request, p *domain.Player) {
	p.SendMsg(r,
		&dto.Msg{
			Action:  dto.CURRENT_ROOM,
			Payload: p.GetRoomId(),
		},
	)
}
