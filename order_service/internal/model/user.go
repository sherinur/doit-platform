package model

import "errors"

type User struct {
	ID    uint
	Email string
}

func (u *User) Validate() error {
	if len(u.Email) == 0 {
		return errors.New("email is required")
	}
	return nil
}
