package handler

import (
	"context"

	"user-services/internal/domain/model"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, request *model.User) (*model.User, error)
	GetUserById(ctx context.Context, userID int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, userID int64) error
	DeleteUser(ctx context.Context, userID int64) error
}
