package server

import "github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend"

type AnswerUseCase interface {
	frontend.AnswerUseCase
}
