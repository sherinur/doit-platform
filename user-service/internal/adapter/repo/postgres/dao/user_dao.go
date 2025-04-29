package dao

import "user-services/internal/domain/model"

type User struct {
	ID       int64  `db:"user_id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func FromDomain(user *model.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func ToDomain(user User) *model.User {
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
