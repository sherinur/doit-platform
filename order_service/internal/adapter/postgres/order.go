package postgres

import (
	"context"
	"database/sql"

	"order_service/internal/adapter/postgres/dao"
	"order_service/internal/model"
)

type Order struct {
	db    *sql.DB
	table string
}

const tableOrders = "orders"

func NewOrder(conn *sql.DB) *Order {
	return &Order{
		db:    conn,
		table: tableOrders,
	}
}

func (o *Order) Create(ctx context.Context, item model.Order) error {
	query := "INSERT INTO " + o.table + " (name, price) VALUES ($1, $2)"
	order := dao.FromOrder(item)

	_, err := o.db.ExecContext(ctx, query, order.Name, order.Price)
	if err != nil {
		return err
	}

	return nil
}

func (o *Order) Get(ctx context.Context, id uint64) (*model.Order, error) {
	query := "SELECT * FROM " + o.table + " WHERE id = $1"
	var order dao.Order

	err := o.db.QueryRowContext(ctx, query, id).Scan(&order)
	if err != nil {
		return nil, err
	}

	item := dao.ToOrder(order)
	return &item, nil
}

func (o *Order) GetAll(ctx context.Context) ([]model.Order, error) {
	return nil, nil
}

func (o *Order) Update(ctx context.Context, order model.Order) error {
	return nil
}

func (o *Order) Delete(ctx context.Context, id uint64) error {
	return nil
}
