package handler

import (
	"context"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, request *model.User) (*model.User, error)
	LoginUser(ctx context.Context, request *model.User) (string, string, error)
	GetUserById(ctx context.Context, userID int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, userID int64) error
	DeleteUser(ctx context.Context, userID int64) error
}
