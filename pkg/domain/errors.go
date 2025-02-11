package domain

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserNotFound    = errors.New("failed get user info")
)
