package playeraction

import (
	"log"
	"net/http"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

func Watch(r *http.Request, msg dto.Msg, p *domain.Player) {
	switch msg.Action {
	case dto.SET_NAME:
		name := msg.Payload.(string)
		p.SetName(name)
		p.SendMsg(r.Context(), dto.Msg{Action: dto.SET_NAME, Payload: name})
		log.Printf("set name %s to %s", name, p.Id())
	default:
	}
}
