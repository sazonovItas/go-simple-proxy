package httphandler

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Message interface{} `json:"message"`
}

func NewError(err error) Error {
	var e Error
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Message = v.Message
	default:
		e.Message = unwrapRecursive(v).Error()
	}

	return e
}

func unwrapRecursive(err error) error {
	originalErr := err

	for originalErr != nil {
		internalErr := errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}

		originalErr = internalErr
	}

	return originalErr
}
