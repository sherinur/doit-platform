package userrepo

import (
	"context"
	"database/sql"
	"user-services/internal/adapter/repo/postgres/dao"
	"user-services/internal/domain/model"
)

type userRepo struct {
	table string
	db    *sql.DB
}

const (
	table_user = "user"
)

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		table: table_user,
		db:    db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	object := dao.FromDomain(user)
	query := `
		INSERT INTO ` + r.table + `(user_id, username, email, password, role)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query, object.ID, object.Username, object.Password, object.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetById(ctx context.Context, user_id int64) (*model.User, error) {
	query := `
        SELECT user_id, username, email, password, role
        FROM ` + r.table + `
        WHERE user_id = $1
    `

	row := r.db.QueryRowContext(ctx, query, user_id)

	user := dao.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return dao.ToDomain(user), err
}

func (r *userRepo) Update(ctx context.Context, user *model.User, user_id int64) error {
	object := dao.FromDomain(user)
	query := `
        UPDATE ` + r.table + `
        SET username = $1, email = $2, password = $3, role = $4
        WHERE user_id = $5
    `
	_, err := r.db.ExecContext(ctx, query, object.Username, object.Email, object.Password, object.Role, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Delete(ctx context.Context, user_id int64) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE user_id = $1
    `
	_, err := r.db.ExecContext(ctx, query, user_id)
	if err != nil {
		return err
	}
	return nil
}
