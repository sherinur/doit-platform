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

type QuestionRepo interface {
	CreateQuestion(ctx context.Context, quiz model.Question) (model.Question, error)
	GetQuestionsByQuizId(ctx context.Context, id string) ([]model.Question, error)
	GetQuestionById(ctx context.Context, id string) (model.Question, error)
	UpdateQuestion(ctx context.Context, quiz model.Question) error
	DeleteQuestion(ctx context.Context, id string) error
}

type AnswerRepo interface {
	CreateAnswer(ctx context.Context, quiz model.Answer) (model.Answer, error)
	GetAnswersByQuestionId(ctx context.Context, id string) ([]model.Answer, error)
	GetAnswersByQuestionIds(ctx context.Context, id []string) ([]model.Answer, error)
	GetAnswerById(ctx context.Context, id string) (model.Answer, error)
	UpdateAnswer(ctx context.Context, quiz model.Answer) error
	DeleteAnswer(ctx context.Context, id string) error
}
