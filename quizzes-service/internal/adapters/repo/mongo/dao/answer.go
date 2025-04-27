package dao

import (
	"quizzes-service/internal/model"
)

type Answer struct {
	ID        string `bson:"_id,omitempty"`
	Text      string `bson:"text"`
	IsCorrect bool   `bson:"is_correct"`
}

func FromAnswer(answer model.Answer) Answer {
	return Answer{
		ID:        answer.ID,
		Text:      answer.Text,
		IsCorrect: answer.IsCorrect,
	}
}

func ToAnswer(answer Answer) model.Answer {
	return model.Answer{
		ID:        answer.ID,
		Text:      answer.Text,
		IsCorrect: answer.IsCorrect,
	}
}
