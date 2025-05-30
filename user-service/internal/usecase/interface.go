package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetById(ctx context.Context, userID int64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	UpdateInfo(ctx context.Context, user *model.UserUpdateData) error
	UpdatePassword(ctx context.Context, user *model.UserUpdateData) error
	Delete(ctx context.Context, userID int64) error
	ChangeUserRole(ctx context.Context, userID int64, newRole string) error
	VerifyEmail(ctx context.Context, email string) error
}

type UserCache interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	SetUser(ctx context.Context, user *model.User) error
	SaveVerificationCode(ctx context.Context, email string, code string) error
	GetVerificationCode(ctx context.Context, email string) (string, error)
	InvalidateUser(ctx context.Context, id int64) error
	InvalidateUsersList(ctx context.Context) error
}

type SessionCache interface {
	SetSession(ctx context.Context, sessionID string, data model.Session) error
	GetSession(ctx context.Context, sessionID string) (*model.Session, error)
	InvalidateSession(ctx context.Context, sessionID string) error
}

type UserEventStorage interface {
	Push(ctx context.Context, client model.User) error
}
