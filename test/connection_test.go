package test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/ibezgin/mobqom-smequiz/config"
	"github.com/ibezgin/mobqom-smequiz/internal/server"
)

func TestConnection(t *testing.T) {
	ctx := context.Background()

	const connCount = 999
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
			conn, _, err := websocket.Dial(ctx, fmt.Sprintf("ws://%s:%s", cfg.Host, cfg.Port), nil)
			defer func() {
				err := conn.Close(websocket.StatusNormalClosure, "клиент завершил работу")
				if err != nil {
					fmt.Println("Ошибка закрытия соединения")
				}
			}()
			log.Println("Клиент подключен к серверу")

			if err != nil {
				log.Printf("Ошибка подключения клиента: %v", err)
				return
			}

			time.Sleep(3 * time.Second)

			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(500 * time.Millisecond) // Give server time to process closures
}
