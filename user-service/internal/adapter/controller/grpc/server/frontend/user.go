package frontend

import (
	"context"

	svc "github.com/sherinur/doit-platform/apis/gen/user-service/service/frontend/user/v1"
	"github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server/frontend/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	svc.UnimplementedUserServiceServer

	uc UserUsecase
}

func NewUser(uc UserUsecase) *User {
	return &User{
		uc: uc,
	}
}

func (u *User) Register(ctx context.Context, req *svc.RegisterRequest) (*svc.RegisterResponse, error) {
	user, err := dto.ToUserFromRegisterRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newuser, err := u.uc.RegisterUser(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromUserToRegisterResponse(*newuser)
}

func (u *User) Login(ctx context.Context, req *svc.LoginRequest) (*svc.LoginResponse, error) {
	user, err := dto.ToUserFromLoginRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accessToken, refreshToken, err := u.uc.LoginUser(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromUserToLoginResponse(accessToken, refreshToken)
}
