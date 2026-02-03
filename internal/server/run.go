package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ibezgin/mobqom-smequiz/config"
	router "github.com/ibezgin/mobqom-smequiz/internal/controller/http"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
)

func Run() {
	cfg := config.InitConfig()
	fmt.Printf("Starting server on port %s\n", cfg.Port)
	m := domain.NewRoomManager()
	go func() {
		for {
			fmt.Println(m)
			time.Sleep(5 * time.Second)
		}
	}()
	router.Init(m)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), nil))
}
