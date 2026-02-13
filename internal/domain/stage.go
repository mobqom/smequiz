package domain

type Stage struct {
	Id       string
	Question *Question
	Players  map[string]*Player
	Answer   map[string]string
}
