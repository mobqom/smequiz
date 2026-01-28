package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type MsgType string

const (
	MsgType_Broadcast MsgType = "broadcast"
)

type Client struct {
	ID     string
	mu     *sync.RWMutex
	conn   *websocket.Conn
	msgCh  chan *ResMsg
	doneCh chan struct{}
}

type ReqMsg struct {
	MsgType MsgType
	Client  *Client
	Data    string
}
type ResMsg struct {
	MsgType  MsgType
	Data     string
	SenderId string
}

func NewResMsg(msg *ReqMsg) *ResMsg {
	return &ResMsg{
		MsgType:  msg.MsgType,
		Data:     msg.Data,
		SenderId: msg.Client.ID,
	}
}

func NewClient(conn *websocket.Conn) *Client {
	ID := rand.Text()
	return &Client{
		ID:     ID,
		mu:     new(sync.RWMutex),
		conn:   conn,
		msgCh:  make(chan *ResMsg, 64),
		doneCh: make(chan struct{}),
	}
}

func (c *Client) readMsgLoop(srv *Server) {
	defer func() {
		close(c.doneCh)
		srv.leaveServerCh <- c
	}()
	for {
		_, b, err := c.conn.ReadMessage()
		if err != nil {

			return
		}
		msg := new(ReqMsg)
		err = json.Unmarshal(b, msg)
		if err != nil {
			fmt.Printf("error unmarshall the msg %v\n", err)
			continue
		}
		msg.Client = c
		srv.broadcastCh <- msg
		// fmt.Println((string)p)
	}
}

func (c *Client) writeMessageLoop() {
	defer c.conn.Close()
	for {
		select {
		case <-c.doneCh:
			return
		case msg := <-c.msgCh:
			err := c.conn.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error sending message to cID = %s", c.ID)
				return
			}

		}
	}
}
