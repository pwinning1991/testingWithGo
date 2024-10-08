package suite

import (
	"errors"
)

type User struct {
	ID    int
	Email string
}

var (
	ErrNotFound   = errors.New("User not found")
	ErrEmailTaken = errors.New("Email already taken")
)

type UserStore interface {
	Create(*User) error
	ByID(id int) (*User, error)
	ByEmail(email string) (*User, error)
	Delete(*User) error
}
