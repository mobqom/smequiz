package dto

type ReqMsg struct {
	Action  MsgType     `json:"action"`
	Payload interface{} `json:"payload"`
}
