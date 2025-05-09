package repo

import (
	"context"
	"database/sql"

	"github.com/sherinur/doit-platform/user-service/internal/adapter/repo/postgres/dao"
	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type userRepo struct {
	table string
	db    *sql.DB
}

const (
	tableUser = "users"
)

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		table: tableUser,
		db:    db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	newuser := dao.FromDomain(user)
	query := `
        INSERT INTO users (name, phone, email, role, password_hash, created_at, updated_at, is_deleted)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
    `

	err := r.db.QueryRowContext(ctx, query,
		newuser.Name,
		newuser.Phone,
		newuser.Email,
		newuser.Role,
		newuser.PasswordHash,
		newuser.CreatedAt,
		newuser.UpdatedAt,
		newuser.IsDeleted,
	).Scan(newuser.ID)
	if err != nil {
		return &model.User{}, err
	}

	return user, err
}

func (r *userRepo) GetById(ctx context.Context, userID int64) (*model.User, error) {
	query := `
        SELECT id, name, phone, email, role, password_hash, created_at, updated_at, is_deleted
        FROM users
        WHERE id = $1 AND is_deleted = false
    `

	row := r.db.QueryRowContext(ctx, query, userID)

	user := dao.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Phone,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)
	if err != nil {
		return nil, err
	}

	return dao.ToDomain(user), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
        SELECT id, name, phone, email, role, password_hash, created_at, updated_at, is_deleted
        FROM users
        WHERE email = $1 AND is_deleted = false
    `

	row := r.db.QueryRowContext(ctx, query, email)

	user := dao.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Phone,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)
	if err != nil {
		return nil, err
	}

	return dao.ToDomain(user), nil
}

func (r *userRepo) GetAll(ctx context.Context) ([]*model.User, error) {
	query := `
        SELECT id, name, phone, email, role, password_hash, created_at, updated_at, is_deleted
        FROM users
        WHERE is_deleted = false
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := dao.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Phone,
			&user.Email,
			&user.Role,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsDeleted,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, dao.ToDomain(user))
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) Update(ctx context.Context, user *model.User, userID int64) error {
	object := dao.FromDomain(user)
	query := `
        UPDATE users
        SET name = $1, phone = $2, email = $3, password_hash = $4, updated_at = $6
        WHERE id = $7 AND is_deleted = false
    `
	_, err := r.db.ExecContext(ctx, query,
		object.Name,
		object.Phone,
		object.Email,
		object.PasswordHash,
		object.UpdatedAt,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Delete(ctx context.Context, userID int64) error {
	query := `
        UPDATE users
        SET is_deleted = true, updated_at = CURRENT_TIMESTAMP
        WHERE id = $1
    `
	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
