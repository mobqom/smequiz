package test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/config"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

func TestJoinRoom(t *testing.T) {
	const connCount = 100
	cfg := &config.AppConfig{
		Host: "localhost",
		Port: "8080",
	}
	go server.Run(cfg)
	time.Sleep(1 * time.Second)
	dialer := websocket.DefaultDialer
	url := fmt.Sprintf("ws://%s:%s/game", cfg.Host, cfg.Port)
	creatorConn, _, err := dialer.Dial(url, nil)
	roomIdCh := make(chan string)
	if err != nil {
		fmt.Printf("error of creator connection\n")

	}
	creatorConn.WriteJSON(dto.Msg{Action: dto.CREATE_ROOM})
	go func() {
		msg := new(dto.Msg)
		_, p, err := creatorConn.ReadMessage()
		if err != nil {
			fmt.Printf("error reading creator message\n")
		}
		err = json.Unmarshal(p, msg)
		if err != nil {
			fmt.Printf("error of unmarshall message %s\n", err)
		}
		for {
			switch msg.Action {
			case dto.CURRENT_ROOM:
				fmt.Println(msg.Payload)
				roomIdCh <- msg.Payload.(string)
				return
			}
		}
	}()
	roomId := <-roomIdCh
	var wg sync.WaitGroup
	for range connCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, _, err := dialer.Dial(url, nil)
			if err != nil {
				fmt.Printf("error connection")
			}
			conn.WriteJSON(dto.Msg{Action: dto.JOIN_ROOM, Payload: roomId})
			time.Sleep(2 * time.Second)
			conn.WriteJSON(dto.Msg{Action: dto.LEAVE_ROOM})
			time.Sleep(2 * time.Second)
		}()

	}
	wg.Wait()
}
