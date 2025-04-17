package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"order_service/internal/adapter/postgres/dao"
	"order_service/internal/model"
)

type Order struct {
	db              *sql.DB
	tableOrders     string
	tableOrderItems string
}

const (
	tableOrders     = "orders"
	tableOrderItems = "order_items"
)

func NewOrder(conn *sql.DB) *Order {
	return &Order{
		db:              conn,
		tableOrders:     tableOrders,
		tableOrderItems: tableOrderItems,
	}
}

func (o *Order) Create(ctx context.Context, order model.Order) (int64, error) {
	ordersQuery := fmt.Sprintf("INSERT INTO %s(status, total_amount) VALUES ($1, $2)", o.tableOrders)

	orderDao := dao.FromOrder(order)

	tx, err := o.db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	res, err := o.db.ExecContext(ctx, ordersQuery, orderDao.Status, orderDao.TotalAmount)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	itemsQuery := fmt.Sprintf("INSERT INTO %s(order_id, product_id, quantity, unit_price) VALUES ($1, $2, $3, $4)", o.tableOrderItems)
	stmt, err := tx.Prepare(itemsQuery)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	for _, item := range order.OrderItems {
		itemDao := dao.FromOrderItem(item)

		_, err := stmt.ExecContext(ctx, itemDao.OrderID, itemDao.ProductID, itemDao.Quantity, itemDao.UnitPrice)
		if err != nil {
			return -1, err
		}
	}

	return id, tx.Commit()
}

func (o *Order) Get(ctx context.Context, id int64) (*model.Order, error) {
	ordersQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", o.tableOrders)
	row := o.db.QueryRowContext(ctx, ordersQuery, id)

	var order dao.Order
	err := row.Scan(&order.ID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	itemsQuery := fmt.Sprintf("SELECT * FROM %s WHERE order_id = $1", o.tableOrderItems)
	rows, err := o.db.QueryContext(ctx, itemsQuery, order.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []dao.OrderItem
	for rows.Next() {
		var item dao.OrderItem
		err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.UnitPrice)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	orderModel := dao.ToOrder(order, items)

	return &orderModel, nil
}

func (o *Order) GetAll(ctx context.Context) ([]model.Order, error) {
	query := fmt.Sprintf(`
		SELECT
			o.id, o.status, o.total_amount, o.created_at, o.updated_at,
			i.id, i.order_id, i.product_id, i.quantity, i.unit_price
		FROM %s o 
		LEFT JOIN %s i ON o.id = i.order_id
		ORDER BY o.id
	`, o.tableOrders, o.tableOrderItems)

	rows, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int64]*model.Order)
	var currentOrderID int64

	for rows.Next() {
		var (
			order dao.Order
			item  dao.OrderItem
		)

		err = rows.Scan(
			&order.ID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt,
			&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.UnitPrice,
		)
		if err != nil {
			return nil, err
		}

		if _, exists := ordersMap[order.ID]; !exists {
			entity := dao.ToOrder(order, nil)
			ordersMap[order.ID] = &entity
			currentOrderID = order.ID
		}

		if item.ID != 0 {
			ordersMap[currentOrderID].OrderItems = append(
				ordersMap[currentOrderID].OrderItems,
				dao.ToOrderItem(item),
			)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	orders := make([]model.Order, 0, len(ordersMap))
	for _, order := range ordersMap {
		orders = append(orders, *order)
	}

	return orders, nil
}

func (o *Order) GetByUser(ctx context.Context) ([]model.Order, error) {
	return nil, nil
}

func (o *Order) Update(ctx context.Context, id int64, updated model.Order) error {
	return nil
}

func (o *Order) Delete(ctx context.Context, id int64) error {
	return nil
}
