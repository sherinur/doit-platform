package dao

import (
	"order_service/internal/model"
	"time"
)

type Order struct {
	ID          int64     `db:"id"`
	Status      string    `db:"status"`
	TotalAmount float64   `db:"total_amount"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func FromOrder(o model.Order) Order {
	return Order{
		ID:          o.ID,
		Status:      o.Status,
		TotalAmount: o.TotalAmount,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}

func ToOrder(o Order, items []OrderItem) model.Order {
	var modelItems []model.OrderItem
	for _, item := range items {
		modelItems = append(modelItems, ToOrderItem(item))
	}

	return model.Order{
		ID:          o.ID,
		Status:      o.Status,
		TotalAmount: o.TotalAmount,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
		OrderItems:  modelItems,
	}
}

type OrderItem struct {
	ID        int64   `db:"id"`
	OrderID   int64   `db:"order_id"`
	ProductID int64   `db:"product_id"`
	Quantity  int     `db:"quantity"`
	UnitPrice float64 `db:"unit_price"`
}

func FromOrderItem(i model.OrderItem) OrderItem {
	return OrderItem{
		ID:        i.ID,
		OrderID:   i.OrderID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		UnitPrice: i.UnitPrice,
	}
}

func ToOrderItem(i OrderItem) model.OrderItem {
	return model.OrderItem{
		ID:        i.ID,
		OrderID:   i.OrderID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		UnitPrice: i.UnitPrice,
	}
}
