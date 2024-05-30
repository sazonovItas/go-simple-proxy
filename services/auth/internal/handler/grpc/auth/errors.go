package auth

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authsvc "github.com/sazonovItas/proxy-manager/services/auth/internal/service/auth"
)

func GRPCError(err error) error {
	unwrapErr := errors.Unwrap(err)

	switch {
	case errors.Is(unwrapErr, authsvc.ErrUserNotFound):
		return status.Errorf(codes.NotFound, "user not found")

	case errors.Is(unwrapErr, authsvc.ErrInvalidCredentials):
		return status.Errorf(codes.PermissionDenied, "invalid password")

	case errors.Is(unwrapErr, authsvc.ErrUserAlreadyExists):
		return status.Errorf(codes.AlreadyExists, "user already exists")

	case errors.Is(unwrapErr, authsvc.ErrUserNotVerified):
		return status.Errorf(codes.PermissionDenied, "user has not verified email")

	default:
		return status.Errorf(codes.Internal, err.Error())
	}
}
