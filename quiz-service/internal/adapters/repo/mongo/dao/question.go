package dao

import (
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type Question struct {
	ID      string   `bson:"_id,omitempty"`
	Text    string   `bson:"text"`
	Type    string   `bson:"type"`
	Points  float64  `bson:"points"`
	QuizID  string   `bson:"quiz_id"`
	Answers []Answer `bson:"answers"`
}

type Answer struct {
	AnswerID  string `bson:"answer_id"`
	Text      string `bson:"text"`
	IsCorrect bool   `bson:"is_correct"`
}

func FromQuestion(question model.Question) Question {
	var result Question
	result.ID = question.ID
	result.Text = question.Text
	result.Type = question.Type
	result.Points = question.Points
	result.QuizID = question.QuizID

	for _, answer := range question.Answers {
		result.Answers = append(result.Answers, Answer{AnswerID: answer.AnswerID, Text: answer.Text, IsCorrect: answer.IsCorrect})
	}

	return result
}

func ToQuestion(question Question) model.Question {
	var result model.Question
	result.ID = question.ID
	result.Text = question.Text
	result.Type = question.Type
	result.Points = question.Points
	result.QuizID = question.QuizID

	for _, answer := range question.Answers {
		result.Answers = append(result.Answers, model.Answer{AnswerID: answer.AnswerID, Text: answer.Text, IsCorrect: answer.IsCorrect})
	}

	return result
}
