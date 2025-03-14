package entity

import "errors"

type User struct {
	ID         uint64 `json:"id"`
	Email      string `json:"email"`
	TelegramID int    `json:"telegramID"`
}

var (
	ErrUserNotFound = errors.New("user not found")
)
