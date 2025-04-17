package model

import (
	"time"
)

type Discount struct {
	ID                 int64
	Name               string
	Description        string
	DiscountPercentage float64
	ApplicableProducts []Product
	StartDate          time.Time
	EndDate            time.Time
	IsActive           bool
}
