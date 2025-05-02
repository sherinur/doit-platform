package frontend

import (
	"context"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
)

type AnswerUseCase interface {
	CreateAnswer(ctx context.Context, request model.Answer) (model.Answer, error)
	CreateAnswers(ctx context.Context, request []model.Answer) ([]model.Answer, error)
	GetAnswerById(ctx context.Context, id string) (model.Answer, error)
	GetAnswersByQuestionId(ctx context.Context, id string) ([]model.Answer, error)
	UpdateAnswer(ctx context.Context, request model.Answer) (model.Answer, error)
	DeleteAnswer(ctx context.Context, id string) (model.Answer, error)
}
