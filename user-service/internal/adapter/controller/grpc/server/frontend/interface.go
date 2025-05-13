package frontend

import (
	"context"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, request *model.User) (*model.User, error)
	LoginUser(ctx context.Context, request *model.User) (model.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (model.Token, error)
	GetUserById(ctx context.Context, userID int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, req *model.UserUpdateData) error
	UpdateUserPassword(ctx context.Context, req *model.UserUpdateData) error
	DeleteUser(ctx context.Context, userID int64) error
}
