package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"order_service/internal/adapter/postgres/dao"
	"order_service/internal/model"
)

type User struct {
	db    *sql.DB
	table string
}

const tableUsers = "users"

func NewUser(conn *sql.DB) *User {
	return &User{
		db:    conn,
		table: tableUsers,
	}
}

func (u *User) Create(ctx context.Context, user model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (email) VALUES ($2)", u.table)

	_, err := u.db.ExecContext(ctx, query, user.Email)
	return err
}

func (u *User) GetAll(ctx context.Context) ([]model.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", u.table)
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user dao.User

		err = rows.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, dao.ToUser(user))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
