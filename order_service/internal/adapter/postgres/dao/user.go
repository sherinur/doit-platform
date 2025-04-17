package dao

import (
	"order_service/internal/model"
	"time"
)

type User struct {
	ID        uint      `db:"id"`
	Email     string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Mapping

func ToUser(u User) model.User {
	return model.User{
		ID:    u.ID,
		Email: u.Email,
	}
}

func FromUser(u model.User) model.User {
	return model.User{
		ID:    u.ID,
		Email: u.Email,
	}
}
