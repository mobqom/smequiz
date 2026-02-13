package dto

type Msg struct {
	Action  ActionType  `json:"action" validate:"required"`
	Payload interface{} `json:"payload"`
}
