package frontend

import (
	"context"

	svc "github.com/sherinur/doit-platform/apis/gen/user-service/service/frontend/user/v1"
	"github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server/frontend/dto"
	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	svc.UnimplementedUserServiceServer

	log *zap.Logger
	uc  UserUsecase
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

	token, err := u.uc.LoginUser(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dto.FromTokenToLoginResponse(token)
}

func (u *User) RefreshToken(
	ctx context.Context, req *svc.RefreshTokenRequest,
) (*svc.RefreshTokenResponse, error) {
	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid refresh token")
	}

	token, err := u.uc.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.RefreshTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (u *User) Profile(ctx context.Context, req *svc.ProfileRequest) (*svc.ProfileResponse, error) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user_id missing in context")
	}

	// Now you can use userID to fetch the user profile
	user, err := u.uc.GetUserById(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get profile")
	}

	return dto.FromUserToProfileResponse(*user)
}

func (u *User) UpdateProfile(ctx context.Context, req *svc.UpdateProfileRequest) (*svc.UpdateProfileResponse, error) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user_id missing in context")
	}

	user := &model.UserUpdateData{
		ID:    userID,
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	err := u.uc.UpdateUserInfo(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.UpdateProfileResponse{
		Status: "OK",
	}, nil
}

func (u *User) UpdatePassword(ctx context.Context, req *svc.UpdatePasswordRequest) (*svc.UpdatePasswordResponse, error) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user_id missing in context")
	}

	user := &model.UserUpdateData{
		ID:          userID,
		NewPassword: req.NewPassword,
	}

	err := u.uc.UpdateUserPassword(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.UpdatePasswordResponse{
		Status: "OK",
	}, nil
}

func (u *User) DeleteAccount(ctx context.Context, req *svc.DeleteAccountRequest) (*svc.DeleteAccountResponse, error) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user_id missing in context")
	}

	err := u.uc.DeleteUser(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.DeleteAccountResponse{
		Status: "OK",
	}, nil
}

func (u *User) GetAllUsers(ctx context.Context, req *svc.GetAllUsersRequest) (*svc.GetAllUsersResponse, error) {
	users, err := u.uc.GetAllUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var respUsers []*svc.User
	for _, user := range users {
		respUsers = append(respUsers, &svc.User{
			Id:    user.ID,
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return &svc.GetAllUsersResponse{
		Users: respUsers,
	}, nil
}

func (u *User) Logout(ctx context.Context, req *svc.LogoutRequest) (*svc.LogoutResponse, error) {
	if req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh token required")
	}

	err := u.uc.Logout(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.LogoutResponse{
		Status: "OK",
	}, nil
}

func (u *User) ChangeUserRole(ctx context.Context, req *svc.ChangeUserRoleRequest) (*svc.ChangeUserRoleResponse, error) {
	if req.UserId == 0 || req.NewRole == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id and new_role are required")
	}

	err := u.uc.ChangeUserRole(ctx, req.UserId, req.NewRole)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &svc.ChangeUserRoleResponse{
		Status: "OK",
	}, nil
}

func (u *User) SendVerificationCode(ctx context.Context, req *svc.SendVerificationCodeRequest) (*svc.SendVerificationCodeResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email required")
	}
	err := u.uc.SendVerificationCode(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &svc.SendVerificationCodeResponse{Status: "OK"}, nil
}

func (u *User) VerifyEmail(ctx context.Context, req *svc.VerifyEmailRequest) (*svc.VerifyEmailResponse, error) {
	if req.Email == "" || req.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "email and code required")
	}
	err := u.uc.VerifyEmail(ctx, req.Email, req.Code)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &svc.VerifyEmailResponse{Status: "OK"}, nil
}
