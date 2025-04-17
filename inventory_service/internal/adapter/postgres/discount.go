package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"inventory_service/internal/adapter/postgres/dao"
	"inventory_service/internal/model"
)

type Discount struct {
	db        *sql.DB
	tableName string
}

const (
	tableDiscounts = "discounts"
)

func NewDiscount(conn *sql.DB) *Discount {
	return &Discount{
		db:        conn,
		tableName: tableDiscounts,
	}
}

func (p *Discount) Create(ctx context.Context, discount model.Discount) error {
	d := dao.FromDiscount(discount)

	query := fmt.Sprintf(`
		INSERT INTO %s (id, name, description, discount_percentage, start_date, end_date, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`, p.tableName)

	_, err := p.db.ExecContext(ctx, query,
		d.ID, d.Name, d.Description, d.DiscountPercentage, d.StartDate, d.EndDate, d.IsActive)
	return err
}

func (p *Discount) GetAllProductsWithPromotion(ctx context.Context) ([]model.Product, error) {
	query := fmt.Sprintf(`
		SELECT p.id, p.name, p.description, p.price, p.category, p.created_at, p.updated_at
		FROM products p
		INNER JOIN product_discounts pd ON p.id = pd.product_id
		INNER JOIN discounts d ON pd.discount_id = d.id
		WHERE d.is_active = true AND CURRENT_DATE BETWEEN d.start_date AND d.end_date`)

	rows, err := p.db.QueryContext(ctx, query)
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

func (p *Discount) Delete(ctx context.Context, discountID int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", p.tableName)

	_, err := p.db.ExecContext(ctx, query, discountID)
	return err
}
