package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

type FilePresenter interface {
	Create(ctx context.Context, file model.File) (string, error)
	Get(ctx context.Context, key string) (*model.File, error)
	Delete(ctx context.Context, key string) error
}

type UserPresenter interface {
	RegisterUser(ctx context.Context, request *model.User) (*model.User, error)
	LoginUser(ctx context.Context, request *model.User) (model.Token, error)
	Logout(ctx context.Context, req string) error
	RefreshToken(ctx context.Context, refreshToken string) (model.Token, error)
	GetUserById(ctx context.Context, userID int64) (*model.User, error)
	UpdateUserInfo(ctx context.Context, req *model.UserUpdateData) error
	UpdateUserPassword(ctx context.Context, req *model.UserUpdateData) error
	DeleteUser(ctx context.Context, userID int64) error
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	ChangeUserRole(ctx context.Context, userID int64, newRole string) error
	SendVerificationCode(ctx context.Context, email string) error
	VerifyEmail(ctx context.Context, email, code string) error
}
