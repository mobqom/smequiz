package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	WSPort = ":3223"
)

type Client struct {
	ID   string
	mu   *sync.RWMutex
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	ID := rand.Text()
	return &Client{
		ID:   ID,
		mu:   new(sync.RWMutex),
		conn: conn,
	}
}

type Server struct {
	clients []*Client
	mu      *sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		clients: []*Client{},
		mu:      new(sync.RWMutex)}
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error on http conn upgrade %v\n", err)
		return
	}
	client := NewClient(conn)
	s.mu.Lock()
	s.clients = append(s.clients, client)
	fmt.Println("clients count:", len(s.clients))
	s.mu.Unlock()
}
func createWSServer() {
	s := NewServer()
	fmt.Printf("Starting ws server on port %s\n", WSPort)
	http.HandleFunc("/", s.handleWs)
	log.Fatal(http.ListenAndServe(WSPort, nil))
}
func main() {
	createWSServer()
}
