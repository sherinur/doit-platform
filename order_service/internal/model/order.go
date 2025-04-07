package model

import "errors"

type Order struct {
	ID    uint64
	Name  string
	Price float64
}

func (o *Order) Validate() error {
	if len(o.Name) == 0 {
		return errors.New("order name is not valid")
	}

	if o.Price < 0 {
		return errors.New("order price is not valid")
	}

	return nil
}
