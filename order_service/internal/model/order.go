package model

import (
	"errors"
	"time"
)

type Order struct {
	ID          int64
	UserID      int64
	Status      string
	TotalAmount float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	OrderItems  []OrderItem
}

type OrderItem struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int
	UnitPrice float64
}

func (o *Order) Validate() error {
	if len(o.OrderItems) == 0 {
		return errors.New("order must contain at least one item")
	}

	validStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"completed":  true,
		"cancelled":  true,
	}

	if !validStatuses[o.Status] {
		return errors.New("invalid order status")
	}

	for _, item := range o.OrderItems {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (i *OrderItem) Validate() error {
	if i.ProductID == 0 {
		return errors.New("product ID is required")
	}

	if i.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if i.UnitPrice < 0 {
		return errors.New("unit price cannot be negative")
	}

	return nil
}
