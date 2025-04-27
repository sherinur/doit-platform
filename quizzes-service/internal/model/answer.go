package model

type Answer struct {
	ID         string
	Text       string
	IsCorrect  bool
	QuestionID string
}
