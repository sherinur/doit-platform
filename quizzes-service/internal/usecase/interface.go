package usecase

import (
	"context"
	"quizzes-service/internal/model"
)

type QuizRepo interface {
	CreateQuiz(ctx context.Context, quiz model.Quiz) (model.Quiz, error)
	GetQuizAll(ctx context.Context) ([]model.Quiz, error)
	GetQuizById(ctx context.Context, id string) (model.Quiz, error)
	UpdateQuiz(ctx context.Context, quiz model.Quiz) error
	DeleteQuiz(ctx context.Context, id string) error
}
