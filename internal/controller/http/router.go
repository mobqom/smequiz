package router

import (
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
)

func Init(m domain.RoomManager) {
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		handleWs(m, w, r)
	})
}
