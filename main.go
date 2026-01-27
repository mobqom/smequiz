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
	mu   *sync.Mutex
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	ID := rand.Text()
	return &Client{
		ID:   ID,
		conn: conn,
		mu:   new(sync.Mutex),
	}
}

type Server struct {
	clients []*Client
	mu      *sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: []*Client{},
		mu:      new(sync.Mutex)}
}

func handleWs(rw http.ResponseWriter, r *http.Request) {}

func main() {

	fmt.Println("Start ws server")
	http.HandleFunc("/", handleWs)
	log.Fatal(http.ListenAndServe(WSPort, nil))
}
