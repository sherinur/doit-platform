package service

import (
	"quizzes-service/internal/adapters/controller/http/service/handler"
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
