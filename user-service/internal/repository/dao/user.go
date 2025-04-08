package dao

import "user-service/internal/model"

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func ToUser(user model.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func FromUser(user User) model.User {
	return model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
