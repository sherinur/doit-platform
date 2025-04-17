package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"inventory_service/internal/adapter/postgres/dao"
	"inventory_service/internal/model"
)

type Product struct {
	db        *sql.DB
	tableName string
}

const (
	tableProducts = "products"
)

func NewProduct(conn *sql.DB) *Product {
	return &Product{
		db:        conn,
		tableName: tableProducts,
	}
}

func (p *Product) Create(ctx context.Context, product model.Product) error {
	query := fmt.Sprintf("INSERT INTO %s(name, description, price, category) VALUES ($1, $2, $3, $4) RETURNING id", p.tableName)
	row := p.db.QueryRowContext(ctx, query, product.Name, product.Description, product.Price, product.Category)

	var id int64
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (p *Product) Get(ctx context.Context, id int64) (*model.Product, error) {
	query := fmt.Sprintf("SELECT id, name, description, price, category, created_at, updated_at FROM %s WHERE id = $1", p.tableName)

	var product dao.Product
	row := p.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return nil, err
	}

	modelProduct := dao.ToProduct(product)
	return &modelProduct, nil
}

func (p *Product) Update(ctx context.Context, id int64, updated model.Product) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET name = $1, description = $2, price = $3, category = $4, updated_at = NOW()
		WHERE id = $5
	`, p.tableName)

	_, err := p.db.ExecContext(ctx, query, updated.Name, updated.Description, updated.Price, updated.Category, id)
	return err
}

func (p *Product) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", p.tableName)
	_, err := p.db.ExecContext(ctx, query, id)
	return err
}

func (p *Product) List(ctx context.Context, category string, limit, offset int) ([]model.Product, error) {
	query := fmt.Sprintf("SELECT id, name, description, price, category, created_at, updated_at FROM %s", p.tableName)
	args := []interface{}{}
	cond := ""

	if category != "" {
		cond += " WHERE category = $1"
		args = append(args, category)
	}

	paginate := fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, limit, offset)

	fullQuery := query + cond + paginate

	rows, err := p.db.QueryContext(ctx, fullQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product dao.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, dao.ToProduct(product))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
