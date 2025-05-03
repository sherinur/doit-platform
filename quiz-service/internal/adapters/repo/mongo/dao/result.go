package dao

import (
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"time"
)

type Result struct {
	ID        string           `bson:"_id,omitempty"`
	UserID    string           `bson:"user_id"`
	QuizID    string           `bson:"quiz_id"`
	Score     float64          `bson:"score"`
	Status    string           `bson:"status"`
	Questions []ResultQuestion `bson:"questions"`
	PassedAt  time.Time        `bson:"passed_at"`
}

type ResultQuestion struct {
	QuestionID string         `bson:"question_id"`
	Answers    []ResultAnswer `bson:"answers"`
}

type ResultAnswer struct {
	AnswerID string `bson:"answer_id"`
}

func FromResult(request model.Result) Result {
	var result Result
	result.ID = request.ID
	result.UserID = request.UserID
	result.QuizID = request.QuizID
	result.Score = request.Score
	result.Status = request.Status
	result.PassedAt = request.PassedAt

	for _, question := range request.Questions {
		q := ResultQuestion{QuestionID: question.ID}
		for _, answer := range question.Answers {
			q.Answers = append(q.Answers, ResultAnswer{AnswerID: answer.AnswerID})
		}
		result.Questions = append(result.Questions, q)
	}

	return result
}

func ToResult(request Result) model.Result {
	var result model.Result
	result.ID = request.ID
	result.UserID = request.UserID
	result.QuizID = request.QuizID
	result.Score = request.Score
	result.PassedAt = request.PassedAt
	for _, question := range request.Questions {
		q := model.Question{ID: question.QuestionID}
		for _, answer := range question.Answers {
			q.Answers = append(q.Answers, model.Answer{AnswerID: answer.AnswerID})
		}
		result.Questions = append(result.Questions, q)
	}

	return result
}
