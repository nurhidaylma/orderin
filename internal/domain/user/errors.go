package user

import "errors"

var (
	ErrNotFound           = errors.New("user not found")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidRole        = errors.New("invalid role")
	ErrWrongPassword      = errors.New("wrong password")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
