package dto

type Msg struct {
	Action  ActionType  `json:"action"`
	Payload interface{} `json:"payload"`
}
