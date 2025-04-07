package dao

import "order_service/internal/model"

type Order struct {
	ID    uint64  `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

func FromOrder(item model.Order) Order {
	return Order{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
}

func ToOrder(item Order) model.Order {
	return model.Order{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
}
