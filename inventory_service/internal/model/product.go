package model

import (
	"errors"
	"strings"
	"time"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Product) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("product name is required")
	}

	if p.Price < 0 {
		return errors.New("product price cannot be negative")
	}

	if strings.TrimSpace(p.Category) == "" {
		return errors.New("product category is required")
	}

	return nil
}
