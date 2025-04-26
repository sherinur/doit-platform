package handler

import (
	"context"
	"quizzes-service/internal/model"
)

type QuizUseCase interface {
	CreateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error)
	GetQuizById(ctx context.Context, id string) (model.Quiz, error)
	GetQuizAll(ctx context.Context) ([]model.Quiz, error)
	UpdateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error)
	DeleteQuiz(ctx context.Context, id string) (model.Quiz, error)
}
