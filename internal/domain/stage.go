package domain

type Stage struct {
	Question *Question
	Players  map[string]*Player
}
