package game

type MsgType string
type Stage string
type Screen string

const (
	MsgType_JoinRoom  MsgType = "join_room"
	MsgType_LeaveRoom MsgType = "leave_room"
)

const (
	Stage_Waiting Stage = "waiting"
	Stage_Active  Stage = "active"
	Stage_Ended   Stage = "ended"
)

const (
	Screen_WaitPlayers Screen = "lobby"
	Screen_Question    Screen = "question"
	Screen_Scoreboard  Screen = "scoreboard"
)

type ReqMsg struct {
	MsgType MsgType     `json:"type"`
	Player  *Player     `json:"player"`
	Data    interface{} `json:"data"`
	RoomID  string      `json:"roomID"`
}
