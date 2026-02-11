package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/config"
	router "github.com/ibezgin/mobqom-smequiz/internal/controller/ws"
	"github.com/ibezgin/mobqom-smequiz/internal/domain"
)

func Run(cfg *config.AppConfig) {
	m := domain.NewRoomManager()

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			router.HandleWebSocket(m, w, r)
		}),
	}
	log.Printf("Сервер запущен на %v", addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("ListenAndServe:", err)
	}
}
