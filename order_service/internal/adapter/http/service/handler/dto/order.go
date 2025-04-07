package dto

import "errors"

type CreateOrder struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (o *CreateOrder) Validate() error {
	if len(o.Name) == 0 {
		return errors.New("order name is not valid")
	}

	if o.Price < 0 {
		return errors.New("order price is not valid")
	}

	return nil
}
