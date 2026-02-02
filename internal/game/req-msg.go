package game

type ReqMsg struct {
	MsgType MsgType     `json:"type"`
	Player  *Player     `json:"player"`
	Data    interface{} `json:"data"`
	RoomID  string      `json:"roomID"`
}
