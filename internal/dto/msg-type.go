package dto

type ActionType string

const (
	CREATE_ROOM  ActionType = "CREATE_ROOM"
	JOIN_ROOM    ActionType = "JOIN_ROOM"
	LEAVE_ROOM   ActionType = "LEAVE_ROOM"
	PLAYERS_LIST ActionType = "PLAYERS_LIST"
	SET_NAME     ActionType = "SET_NAME"
	CURRENT_ROOM ActionType = "CURRENT_ROOM"
	START_GAME   ActionType = "START_GAME"
	SET_SCREEN   ActionType = "SET_SCREEN"
	TIMER_TIME   ActionType = "TIMER_TIME"
	SET_QUESTION ActionType = "SET_QUESTION"
)
