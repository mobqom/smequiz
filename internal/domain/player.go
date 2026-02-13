package domain

import (
	"context"
	"log"
	"sync"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

type Player struct {
	conn   *websocket.Conn
	id     string
	name   string
	roomId string
	mu     sync.Mutex
}

func NewPlayer(conn *websocket.Conn, id string) *Player {
	return &Player{
		conn: conn,
		id:   id,
		name: "",
	}
}
func (p *Player) SetRoomId(roomId string) {
	p.roomId = roomId
}
func (p *Player) GetRoomId() string {
	return p.roomId
}
func (p *Player) SetName(name string) {
	p.name = name
}
func (p *Player) GetConn() *websocket.Conn {
	return p.conn
}
func (p *Player) GetId() string {
	return p.id
}
func (p *Player) SendMsg(ctx context.Context, msg dto.Msg) {
	err := wsjson.Write(ctx, p.GetConn(), msg)
	if err != nil {

		log.Printf("%s: send msg err: %v", p.GetId(), err)
	}
}
