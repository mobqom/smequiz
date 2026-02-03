package domain

import (
	"github.com/gorilla/websocket"
)

type Player interface {
	SetRoomId(roomId string)
	GetRoomId() string
	SetName(name string)
	GetConn() *websocket.Conn
	GetId() string
}
type player struct {
	Conn   *websocket.Conn
	Id     string
	Name   string
	RoomId string
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
