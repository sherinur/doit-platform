package dao

import (
	"time"

	"inventory_service/internal/model"
)

type Discount struct {
	ID                 int64     `db:"id"`
	Name               string    `db:"name"`
	Description        string    `db:"description"`
	DiscountPercentage float64   `db:"discount_percentage"`
	StartDate          time.Time `db:"start_date"`
	EndDate            time.Time `db:"end_date"`
	IsActive           bool      `db:"is_active"`
}

func FromDiscount(d model.Discount) Discount {
	return Discount{
		ID:                 d.ID,
		Name:               d.Name,
		Description:        d.Description,
		DiscountPercentage: d.DiscountPercentage,
		StartDate:          d.StartDate,
		EndDate:            d.EndDate,
		IsActive:           d.IsActive,
	}
}

func ToDiscount(d Discount, products []model.Product) model.Discount {
	return model.Discount{
		ID:                 d.ID,
		Name:               d.Name,
		Description:        d.Description,
		DiscountPercentage: d.DiscountPercentage,
		ApplicableProducts: products,
		StartDate:          d.StartDate,
		EndDate:            d.EndDate,
		IsActive:           d.IsActive,
	}
}
