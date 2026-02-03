package domain

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	Id   string
	Name string
}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		conn: conn,
	}
}
