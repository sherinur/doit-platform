package handler

import (
	"context"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type QuizUseCase interface {
	CreateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error)
	GetQuizById(ctx context.Context, id string) (model.Quiz, error)
	UpdateQuiz(ctx context.Context, request model.Quiz) (model.Quiz, error)
	DeleteQuiz(ctx context.Context, id string) (model.Quiz, error)
}

type QuestionUseCase interface {
	CreateQuestion(ctx context.Context, request model.Question) (model.Question, error)
	CreateQuestions(ctx context.Context, request []model.Question) ([]model.Question, error)
	GetQuestionById(ctx context.Context, id string) (model.Question, error)
	GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error)
	UpdateQuestion(ctx context.Context, request model.Question) (model.Question, error)
	DeleteQuestion(ctx context.Context, id string) (model.Question, error)
}

type ResultUseCase interface {
	CreateResult(ctx context.Context, request model.Result) (model.Result, error)
	GetResultById(ctx context.Context, id string) (model.Result, error)
	GetResultsByQuizId(ctx context.Context, id string) ([]model.Result, error)
	GetResultsByUserId(ctx context.Context, id string) ([]model.Result, error)
	DeleteResult(ctx context.Context, id string) (model.Result, error)
}
