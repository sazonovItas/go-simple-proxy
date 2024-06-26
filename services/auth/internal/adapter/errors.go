package adapter

import "errors"

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrUserWithEmailAlreadyExists = errors.New("user with that email already exists")
	ErrUserWithLoginAlreadyExists = errors.New("user with that login already exists")

	ErrVerifyTokenNotFound = errors.New("verify token not found")
	ErrResetTokenNotFound  = errors.New("reset token not found")
)
