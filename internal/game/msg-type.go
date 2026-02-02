package game

type MsgType string

const (
	MsgType_JoinRoom  MsgType = "join_room"
	MsgType_LeaveRoom MsgType = "leave_room"
)
