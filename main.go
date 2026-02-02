package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/game"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

const (
	Port = ":8080"
)

func main() {

	fmt.Printf("Starting WebSocket server on port %s\n", Port)
	gm := game.NewGameManager()
	s := server.NewWSServer()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.HandleWs(gm, w, r)
	})
	http.HandleFunc("/api/create-room", func(w http.ResponseWriter, r *http.Request) {
		s.HandleCreateRoom(gm, w, r)
	})
	go s.AcceptLoop()

	log.Fatal(http.ListenAndServe(Port, nil))

}
