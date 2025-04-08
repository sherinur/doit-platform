package repository

import "user-service/internal/domain"

type UserRepository interface {
	Save(user domain.User) error
}