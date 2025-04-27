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

type QuestionUseCase interface {
	CreateQuestion(ctx context.Context, request model.Question) (model.Question, error)
	GetQuestionById(ctx context.Context, id string) (model.Question, error)
	GetQuestionAll(ctx context.Context) ([]model.Question, error)
	UpdateQuestion(ctx context.Context, request model.Question) (model.Question, error)
	DeleteQuestion(ctx context.Context, id string) (model.Question, error)
}

type AnswerUseCase interface {
	CreateAnswer(ctx context.Context, request model.Answer) (model.Answer, error)
	GetAnswerById(ctx context.Context, id string) (model.Answer, error)
	GetAnswerAll(ctx context.Context) ([]model.Answer, error)
	UpdateAnswer(ctx context.Context, request model.Answer) (model.Answer, error)
	DeleteAnswer(ctx context.Context, id string) (model.Answer, error)
}
