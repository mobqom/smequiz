package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/config"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

func TestConnection(t *testing.T) {
	const connCount = 2
	cfg := &config.AppConfig{
		Host: "localhost",
		Port: "8080",
	}
	go server.Run(cfg)
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	for range connCount {
		wg.Add(1)
		go func() {
			dialer := websocket.DefaultDialer
			conn, _, err := dialer.Dial(
				fmt.Sprintf("ws://%s:%s/game-actions", cfg.Host, cfg.Port), nil)
			if err != nil {
				t.Logf("Connection error: %v", err)
				return
			}

			time.Sleep(3 * time.Second)

			// send a proper close control message so server sees a normal closure
			_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			conn.Close()
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(500 * time.Millisecond) // Give server time to process closures
}
