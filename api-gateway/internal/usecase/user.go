package usecase

import (
	"context"

	"github.com/sherinur/doit-platform/api-gateway/internal/model"
)

type User struct {
	UserPresenter UserPresenter
}

func NewUser(presenter UserPresenter) *User {
	return &User{
		UserPresenter: presenter,
	}
}

func (c *User) RegisterUser(ctx context.Context, request *model.User) (*model.User, error) {
	User, err := c.UserPresenter.RegisterUser(ctx, request)
	if err != nil {
		return &model.User{}, err
	}

	return User, nil
}

func (c *User) LoginUser(ctx context.Context, request *model.User) (model.Token, error) {
	token, err := c.UserPresenter.LoginUser(ctx, request)
	if err != nil {
		return model.Token{}, err
	}

	return token, nil
}

func (c *User) Logout(ctx context.Context, req string) error {
	err := c.UserPresenter.Logout(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	token, err := c.UserPresenter.RefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}

	return token, nil
}

func (c *User) GetUserById(ctx context.Context, userID int64) (*model.User, error) {
	user, err := c.UserPresenter.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *User) UpdateUserInfo(ctx context.Context, req *model.UserUpdateData) error {
	err := c.UserPresenter.UpdateUserInfo(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) UpdateUserPassword(ctx context.Context, req *model.UserUpdateData) error {
	err := c.UserPresenter.UpdateUserPassword(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) DeleteUser(ctx context.Context, userID int64) error {
	err := c.UserPresenter.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := c.UserPresenter.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *User) ChangeUserRole(ctx context.Context, userID int64, newRole string) error {
	err := c.UserPresenter.ChangeUserRole(ctx, userID, newRole)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) SendVerificationCode(ctx context.Context, email string) error {
	err := c.UserPresenter.SendVerificationCode(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) VerifyEmail(ctx context.Context, email, code string) error {
	err := c.UserPresenter.VerifyEmail(ctx, email, code)
	if err != nil {
		return err
	}

	return nil
}
