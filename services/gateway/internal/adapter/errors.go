package adapter

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrUserInvalidCreditanals = errors.New("user invalid creditanals")

	ErrRequestNotFound = errors.New("requests not found")
	ErrInvalidArgument = errors.New("invalid argument")
)
