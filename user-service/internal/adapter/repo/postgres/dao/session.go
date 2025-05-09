package dao

import (
	"time"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type Session struct {
	UserID       int64     `bson:"userID"`
	RefreshToken string    `bson:"refreshToken"`
	ExpiresAt    time.Time `bson:"expiresAt"`
	CreatedAt    time.Time `bson:"createdAt"`
}

func FromSession(session model.Session) Session {
	return Session{
		UserID:       session.UserID,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}
}

func ToSession(daoSession Session) model.Session {
	return model.Session{
		UserID:       daoSession.UserID,
		RefreshToken: daoSession.RefreshToken,
		ExpiresAt:    daoSession.ExpiresAt,
		CreatedAt:    daoSession.CreatedAt,
	}
}
