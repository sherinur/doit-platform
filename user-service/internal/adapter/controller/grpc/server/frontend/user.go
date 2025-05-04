package frontend

import (
	"context"

	usersvc "github.com/sherinur/doit-platform/apis/gen/user-service/service/frontend/user/v1"
)

type User struct {
	uc UserUsecase
}

func NewUser(uc UserUsecase) *User {
	return &User{
		uc: uc,
	}
}

func (u *User) Register(ctx context.Context, req *usersvc.RegisterRequest) (*usersvc.RegisterResponse, error) {
}
