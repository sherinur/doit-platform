package dao

import (
	"quizzes-service/internal/model"
)

type Question struct {
	ID        string   `bson:"_id,omitempty"`
	Text      string   `bson:"text"`
	Type      string   `bson:"type"`
	AnswerIDs []string `bson:"answer_ids"`
}

func FromQuestion(question model.Question) Question {
	return Question{
		ID:        question.ID,
		Text:      question.Text,
		Type:      question.Type,
		AnswerIDs: question.AnswerIDs,
	}
}

func ToQuestion(question Question) model.Question {
	return model.Question{
		ID:        question.ID,
		Text:      question.Text,
		Type:      question.Type,
		AnswerIDs: question.AnswerIDs,
	}
}
