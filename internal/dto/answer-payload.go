package dto

type AnswerPayload struct {
	StageId string `json:"stageId" validate:"required"`
	Answer  string `json:"answer" validate:"required"`
}
