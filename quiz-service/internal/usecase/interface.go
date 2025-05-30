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

type QuizRedis interface {
	SetQuiz(ctx context.Context, key string, quiz model.Quiz) error
	GetQuiz(ctx context.Context, key string) (model.Quiz, error)
}

type QuestionRepo interface {
	CreateQuestion(ctx context.Context, question model.Question) (model.Question, error)
	CreateQuestions(ctx context.Context, question []model.Question) ([]model.Question, error)
	GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error)
	GetQuestionById(ctx context.Context, id string) (model.Question, error)
	UpdateQuestion(ctx context.Context, question model.Question) error
	DeleteQuestion(ctx context.Context, id string) error
}

type QuestionRedis interface {
	SetQuestion(ctx context.Context, key string, question model.Question) error
	GetQuestion(ctx context.Context, key string) (model.Question, error)
	SetManyQuestion(ctx context.Context, questions []model.Question) error
	GetManyQuestion(ctx context.Context, keys []string) ([]model.Question, error)
}

type ResultRepo interface {
	CreateResult(ctx context.Context, result model.Result) (model.Result, error)
	GetResultById(ctx context.Context, id string) (model.Result, error)
	GetResultsByQuizId(ctx context.Context, id string) ([]model.Result, error)
	GetResultsByUserId(ctx context.Context, id string) ([]model.Result, error)
	DeleteResult(ctx context.Context, id string) error
}

type ResultRedis interface {
	SetResult(ctx context.Context, key string, result model.Result) error
	GetResult(ctx context.Context, key string) (model.Result, error)
}
