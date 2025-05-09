package usecase

import (
	"context"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type QuizRepo interface {
	CreateQuiz(ctx context.Context, quiz model.Quiz) (model.Quiz, error)
	GetQuizById(ctx context.Context, id string) (model.Quiz, error)
	UpdateQuiz(ctx context.Context, quiz model.Quiz) error
	ChangeTotalPointsQuiz(ctx context.Context, id string, change float64) error
	DeleteQuiz(ctx context.Context, id string) error
}

type QuestionRepo interface {
	CreateQuestion(ctx context.Context, question model.Question) (model.Question, error)
	CreateQuestions(ctx context.Context, question []model.Question) ([]model.Question, error)
	GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error)
	GetQuestionById(ctx context.Context, id string) (model.Question, error)
	UpdateQuestion(ctx context.Context, question model.Question) error
	DeleteQuestion(ctx context.Context, id string) error
}

type ResultRepo interface {
	CreateResult(ctx context.Context, result model.Result) (model.Result, error)
	GetResultById(ctx context.Context, id string) (model.Result, error)
	GetResultsByQuizId(ctx context.Context, id string) ([]model.Result, error)
	GetResultsByUserId(ctx context.Context, id string) ([]model.Result, error)
	DeleteResult(ctx context.Context, id string) error
}
