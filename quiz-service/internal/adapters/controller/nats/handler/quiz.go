package handler

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/nats/handler/dto"
)

type Quiz struct {
	usecase QuizUsecase
}

func NewQuiz(usecase QuizUsecase) *Quiz {
	return &Quiz{usecase: usecase}
}

func (c *Quiz) Handler(ctx context.Context, msg *nats.Msg) error {
	id, err := dto.ToId(msg)
	if err != nil {
		log.Println("failed to convert Course NATS msg", err)

		return err
	}

	_, err = c.usecase.DeleteQuiz(ctx, id)
	if err != nil {
		log.Println("failed to delete Quiz", err)

		return err
	}

	return nil
}
