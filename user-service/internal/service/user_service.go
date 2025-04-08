package usecase

import "user-service/internal/domain"

type UserUsecase interface {
	CreateUser(user domain.User) error
}