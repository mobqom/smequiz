package test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/ibezgin/mobqom-smequiz/config"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

func TestJoinRoom(t *testing.T) {
	ctx := context.Background()
	const connCount = 100
	cfg := &config.AppConfig{
		Host: "localhost",
		Port: "8080",
	}
	go server.Run(cfg)
	time.Sleep(1 * time.Second)
	url := fmt.Sprintf("ws://%s:%s/game", cfg.Host, cfg.Port)
	creatorConn, _, err := websocket.Dial(ctx, url, nil)
	roomIdCh := make(chan string)
	if err != nil {
		log.Printf("error of creator connection\n")

	}
	wsjson.Write(ctx, creatorConn, dto.Msg{Action: dto.CREATE_ROOM})
	go func() {
		msg := new(dto.Msg)
		err := wsjson.Read(ctx, creatorConn, msg)
		if err != nil {
			log.Printf("error reading creator message\n")
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
	fmt.Println(roomId)
	var wg sync.WaitGroup
	for range connCount {
		wg.Add(1)
		go func() {
			conn, _, err := websocket.Dial(ctx, url, nil)
			defer func() {
				err := wsjson.Write(ctx, conn, dto.Msg{Action: dto.LEAVE_ROOM})
				if err != nil {
					log.Printf("error leaving room")
				}
				wg.Done()
			}()
			if err != nil {
				log.Printf("error connection")
			}
			wsjson.Write(ctx, conn, dto.Msg{Action: dto.JOIN_ROOM, Payload: roomId})
			time.Sleep(1 * time.Second)

		}()

	}
	wg.Wait()
}
