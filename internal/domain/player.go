package domain

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

type Player interface {
	SetRoomId(roomId string)
	GetRoomId() string
	SetName(name string)
	GetConn() *websocket.Conn
	GetId() string
	SendMsg(msg *dto.Msg)
}
type player struct {
	Conn   *websocket.Conn
	Id     string
	Name   string
	RoomId string
	mu     sync.Mutex
}

func NewPlayer(conn *websocket.Conn, id string) Player {
	return &player{
		Conn: conn,
		Id:   id,
		Name: "",
	}
}
func (p *player) SetRoomId(roomId string) {
	p.RoomId = roomId
}
func (p *player) GetRoomId() string {
	return p.RoomId
}
func (p *player) SetName(name string) {
	p.Name = name
}
func (p *player) GetConn() *websocket.Conn {
	return p.Conn
}
func (p *player) GetId() string {
	return p.Id
}
func (p *player) SendMsg(msg *dto.Msg) {
	p.mu.Lock()
	err := p.GetConn().WriteJSON(msg)
	p.mu.Unlock()
	if err != nil {
		fmt.Printf("%s: send msg err: %v\n", p.GetId(), err)
	}
}
