package handler

import (
	"context"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type QuizUsecase interface {
	DeleteQuiz(ctx context.Context, id string) (model.Quiz, error)
}
