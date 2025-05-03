package server

import "github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend"

type QuestionUseCase interface {
	frontend.QuestionUseCase
}

type QuizUseCase interface {
	frontend.QuizUseCase
}

type ResultUseCase interface {
	frontend.ResultUseCase
}
