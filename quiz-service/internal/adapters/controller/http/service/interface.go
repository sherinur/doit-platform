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

type AnswerUseCase interface {
	handler.AnswerUseCase
}

type ResultUseCase interface {
	handler.ResultUseCase
}
