package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSServer struct {
	clients       map[string]*Client
	joinServerCh  chan *Client
	leaveServerCh chan *Client
}

func NewWSServer() *WSServer {
	return &WSServer{
		clients:       map[string]*Client{},
		joinServerCh:  make(chan *Client),
		leaveServerCh: make(chan *Client),
	}
}

func (s *WSServer) HandleWs(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error upgrade websocket connection\n")
	}
	client := NewClient(conn)
	s.joinServerCh <- client
	go client.readMsg()
}

func (s *WSServer) Run(port string) {
	fmt.Printf("Starting WebSocket server on port %s\n", port)
	http.HandleFunc("/ws", s.HandleWs)
}

func (s *WSServer) joinServer(client *Client) {
	s.clients[client.ID] = client
	fmt.Printf("Client %s joined the server\n", client.ID)
}
func (s *WSServer) leaveServer(client *Client) {
	delete(s.clients, client.ID)
	fmt.Printf("Client %s left the server\n", client.ID)
}
