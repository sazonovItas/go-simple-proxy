package authsvc

import (
	"errors"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
)

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrUserWithEmailAlreadyExists = errors.New("user with this email already exists")
	ErrUserWithLoginAlreadyExists = errors.New("user with this login already exists")
	ErrUserEmailNotVerified       = errors.New("user has not verified email")
	ErrInvalidCredentials         = errors.New("invalid password")

	ErrTokenExpired   = errors.New("token expired")
	ErrTokenMalformed = errors.New("token malformed")
)

func AuthErrors(err error) error {
	switch {
	case errors.Is(err, adapter.ErrUserNotFound):
		return ErrUserNotFound

	case errors.Is(err, adapter.ErrUserWithEmailAlreadyExists):
		return ErrUserWithEmailAlreadyExists

	case errors.Is(err, adapter.ErrUserWithLoginAlreadyExists):
		return ErrUserWithLoginAlreadyExists
	}

	return err
}
