package dao

import "github.com/sherinur/doit-platform/quiz-service/internal/model"

type Question struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Points  float64  `json:"points"`
	QuizID  string   `json:"quiz_id"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	AnswerID  string `json:"answer_id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
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
