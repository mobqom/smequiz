package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type TestConfig struct {
	host string
	port string
}

func DialSever() *websocket.Conn {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(fmt.Sprintf("%s%s", "ws://localhost", Port+"/ws"), nil)
	if err != nil {
		fmt.Printf("error with dealer %s\n", err)
	}
	return conn
}
func TestConnection(t *testing.T) {
	go startApp()
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	for range 500 {
		wg.Add(1)
		go func() {
			conn := DialSever()
			time.Sleep(1 * time.Second)
			defer func() {
				conn.Close()
				wg.Done()
			}()
		}()
	}
	wg.Wait()
}
