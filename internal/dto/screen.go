package dto

type Screen string

const (
	WAITING_SCREEN         Screen = "WAITING_SCREEN"
	TIMER_SCREEN           Screen = "TIMER_SCREEN"
	QUESTION_SCREEN        Screen = "QUESTION_SCREEN"
	QUESTION_RESULT_SCREEN Screen = "QUESTION_RESULT_SCREEN"
	MESSAGE_SCREEN         Screen = "MESSAGE_SCREEN"
)
