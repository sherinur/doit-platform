package dao

import (
	"quizzes-service/internal/model"
	"time"
)

type Result struct {
	ID       string    `bson:"_id,omitempty"`
	UserID   string    `bson:"user_id"`
	QuizID   string    `bson:"quiz_id"`
	Score    float64   `bson:"score"`
	PassedAt time.Time `bson:"passed_at"`
}

func FromResult(result model.Result) Result {
	return Result{
		ID:       result.ID,
		UserID:   result.UserID,
		QuizID:   result.QuizID,
		Score:    result.Score,
		PassedAt: result.PassedAt,
	}
}

func ToResult(result Result) model.Result {
	return model.Result{
		ID:       result.ID,
		UserID:   result.UserID,
		QuizID:   result.QuizID,
		Score:    result.Score,
		PassedAt: result.PassedAt,
	}
}
