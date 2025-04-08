package repository

import (
	"context"
	"database/sql"
	"user-service/internal/model"
	"user-service/internal/repository/dao"
)

type userRepository struct {
	db    *sql.DB
	table string
}

const (
	tabeleUser = "User"
)

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db:    db,
		table: tabeleUser,
	}
}

func (r *userRepository) Create(ctx context.Context, user model.User) error {
	object := dao.ToUser(user)
	query := `
	INSERT INTO ` + r.table + ` (id, username, email, password, role)
	VALUES $1, $2, $3, $4, $5`

	_, err := r.db.ExecContext(ctx, query, object.ID, object.Username, object.Email, object.Password, object.Role)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (model.User, error) {
	query := `
	SELECT id, name, description
	FROM ` + r.table + `
	WHERE id = $1
`
	row := r.db.QueryRowContext(ctx, query, id)

	var user dao.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return model.User{}, err
	}

	return dao.FromUser(user), nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `
        DELETE FROM ` + r.table + `
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
