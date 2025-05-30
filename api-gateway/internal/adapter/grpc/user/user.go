package user

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
	svc "github.com/sherinur/doit-platform/apis/gen/user-service/service/frontend/user/v1"
)

type User struct {
	user svc.UserServiceClient
}

func NewUser(user svc.UserServiceClient) *User {
	return &User{
		user: user,
	}
}

func (c *User) RegisterUser(ctx context.Context, request *model.User) (*model.User, error) {
	resp, err := c.user.Register(ctx, &svc.RegisterRequest{
		Email:    request.Email,
		Password: request.CurrentPassword,
	})

	if err != nil {
		return &model.User{}, err
	}

	user := model.User{
		ID: resp.UserId,
	}

	return &user, err
}

func (c *User) LoginUser(ctx context.Context, request *model.User) (model.Token, error) {
	resp, err := c.user.Login(ctx, &svc.LoginRequest{
		Email:    request.Email,
		Password: request.CurrentPassword,
	})

	if err != nil {
		return model.Token{}, err
	}

	token := model.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}

	return token, nil
}

func (c *User) Logout(ctx context.Context, req string) error {
	_, err := c.user.Logout(ctx, &svc.LogoutRequest{
		RefreshToken: req,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *User) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	resp, err := c.user.RefreshToken(ctx, &svc.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})

	if err != nil {
		return model.Token{}, err
	}

	token := model.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}

	return token, nil
}

func (c *User) GetUserById(ctx context.Context, userID int64) (*model.User, error) {
	resp, err := c.user.Profile(ctx, &svc.ProfileRequest{})

	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:    resp.Id,
		Name:  resp.Name,
		Email: resp.Email,
		Phone: resp.Phone,
	}

	return user, nil
}

func (c *User) UpdateUserInfo(ctx context.Context, req *model.UserUpdateData) error {
	_, err := c.user.UpdateProfile(ctx, &svc.UpdateProfileRequest{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *User) UpdateUserPassword(ctx context.Context, req *model.UserUpdateData) error {
	_, err := c.user.UpdatePassword(ctx, &svc.UpdatePasswordRequest{
		CurrentPassword: req.NewPassword,
		NewPassword:     req.NewPassword,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *User) DeleteUser(ctx context.Context, userID int64) error {
	_, err := c.user.DeleteAccount(ctx, &svc.DeleteAccountRequest{})

	if err != nil {
		return err
	}

	return nil
}

func (c *User) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	resp, err := c.user.GetAllUsers(ctx, &svc.GetAllUsersRequest{})

	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, u := range resp.Users {
		users = append(users, &model.User{
			ID:    u.Id,
			Name:  u.Name,
			Email: u.Email,
			Phone: u.Phone,
		})
	}

	return users, nil
}

func (c *User) ChangeUserRole(ctx context.Context, userID int64, newRole string) error {
	_, err := c.user.ChangeUserRole(ctx, &svc.ChangeUserRoleRequest{
		UserId:  userID,
		NewRole: newRole,
	})

	if err != nil {
		return err
	}

	return nil
}
func (c *User) SendVerificationCode(ctx context.Context, email string) error {
	_, err := c.user.SendVerificationCode(ctx, &svc.SendVerificationCodeRequest{
		Email: email,
	})

	if err != nil {
		return err
	}

	return nil
}
func (c *User) VerifyEmail(ctx context.Context, email, code string) error {
	_, err := c.user.VerifyEmail(ctx, &svc.VerifyEmailRequest{
		Email: email,
		Code:  code,
	})

	if err != nil {
		return err
	}

	return nil
}
