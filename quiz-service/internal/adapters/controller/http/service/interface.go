package service

import (
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/http/service/handler"
)

type QuizUseCase interface {
	handler.QuizUseCase
}

type QuestionUseCase interface {
	handler.QuestionUseCase
}

type ResultUseCase interface {
	handler.ResultUseCase
}
