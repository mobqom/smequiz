package server

import (
	"crypto/rand"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	ID := rand.Text()[:9]

	return &Client{
		ID:   ID,
		Conn: conn,
	}
}

func (c *Client) readMsg() {
	defer c.Conn.Close()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("error reading message: %v\n", err)
			break
		}
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}

}
