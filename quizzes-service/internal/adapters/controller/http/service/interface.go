package service

import (
	"quizzes-service/internal/adapters/controller/http/service/handler"
)

type QuizUseCase interface {
	handler.QuizUseCase
}
