package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sherinur/doit-platform/user-service/internal/adapter/repo/postgres/dao"
	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type sessionRepo struct {
	table string
	db    *sql.DB
}

const (
	tableSession = "users.sessions"
)

func NewSessionRepo(db *sql.DB) *sessionRepo {
	return &sessionRepo{
		table: tableSession,
		db:    db,
	}
}

func (r *sessionRepo) Create(ctx context.Context, session *model.Session) error {
	query := `
        INSERT INTO ` + r.table + ` (user_id, refresh_token, expires_at, created_at)
        VALUES ($1, $2, $3, $4)
    `

	_, err := r.db.ExecContext(ctx, query, session.UserID, session.RefreshToken, session.ExpiresAt, session.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}

func (r *sessionRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*model.Session, error) {
	query := `
        SELECT user_id, refresh_token, expires_at, created_at
        FROM ` + r.table + `
        WHERE refresh_token = $1
    `

	row := r.db.QueryRowContext(ctx, query, refreshToken)

	var daoSession dao.Session
	err := row.Scan(&daoSession.UserID, &daoSession.RefreshToken, &daoSession.ExpiresAt, &daoSession.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get session by refresh token: %w", err)
	}

	session := dao.ToSession(daoSession)
	return &session, nil
}

func (r *sessionRepo) DeleteByRefreshToken(ctx context.Context, refreshToken string) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE refresh_token = $1
    `

	_, err := r.db.ExecContext(ctx, query, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to delete session by refresh token: %w", err)
	}

	return nil
}

func (r *sessionRepo) DeleteByUserID(ctx context.Context, userID uint64) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE user_id = $1
    `

	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete sessions by user ID: %w", err)
	}

	return nil
}
