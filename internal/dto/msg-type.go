package dto

type MsgType string

const (
	CREATE_ROOM MsgType = "CREATE_ROOM"
	JOIN_ROOM   MsgType = "JOIN_ROOM"
	LEAVE_ROOM  MsgType = "LEAVE_ROOM"
)
