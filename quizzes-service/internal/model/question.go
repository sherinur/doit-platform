package model

type Question struct {
	ID        string
	Text      string
	Type      string
	AnswerIDs []string
}
