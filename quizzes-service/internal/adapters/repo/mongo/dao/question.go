package dao

import (
	"quizzes-service/internal/model"
)

type Question struct {
	ID     string  `bson:"_id,omitempty"`
	Text   string  `bson:"text"`
	Type   string  `bson:"type"`
	Points float64 `bson:"points"`
	QuizID string  `bson:"quiz_id"`
}

func FromQuestion(question model.Question) Question {
	return Question{
		ID:     question.ID,
		Text:   question.Text,
		Type:   question.Type,
		Points: question.Points,
		QuizID: question.QuizID,
	}
}

func ToQuestion(question Question) model.Question {
	return model.Question{
		ID:     question.ID,
		Text:   question.Text,
		Type:   question.Type,
		Points: question.Points,
		QuizID: question.QuizID,
	}
}
