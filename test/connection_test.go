package test

import (
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

func TestConnection(t *testing.T) {
	const connCount = 999
	go server.Run()
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	for range connCount {
		wg.Add(1)
		go func() {
			dialer := websocket.DefaultDialer
			conn, _, err := dialer.Dial("ws://localhost:8080/game", nil)
			if err != nil {
				t.Logf("Connection error: %v", err)
				return
			}

			time.Sleep(3 * time.Second)

			defer func() {
				// send a proper close control message so server sees a normal closure
				_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				conn.Close()
				wg.Done()
			}()
		}()
	}
	wg.Wait()
}
