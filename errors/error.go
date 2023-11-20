package myerrors

import "errors"

var (
	ErrInvalidUserID = errors.New("invalid user ID")
	ErrUserNotFound  = errors.New("user not found")
	// Adicione outros erros conforme necessário
)
