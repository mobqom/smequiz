package playeraction

import (
	"fmt"

	"github.com/ibezgin/mobqom-smequiz/internal/domain"
	"github.com/ibezgin/mobqom-smequiz/internal/dto"
)

func Watch(msg *dto.Msg, p *domain.Player) {
	switch msg.Action {
	case dto.SET_NAME:
		name := msg.Payload.(string)
		p.SetName(name)
		p.SendMsg(&dto.Msg{Action: dto.SET_NAME, Payload: name})
		fmt.Printf("set name %s to %s\n", name, p.GetId())
	default:
	}
}
