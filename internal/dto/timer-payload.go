package dto

type TimerPayload struct {
	Value int  `json:"value"`
	Done  bool `json:"done"`
}
