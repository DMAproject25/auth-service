package entity

import "errors"

type User struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

var (
	ErrUserNotFound = errors.New("user not found")
)
