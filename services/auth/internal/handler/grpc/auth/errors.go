package auth

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authsvc "github.com/sazonovItas/proxy-manager/services/auth/internal/service/auth"
)

func GRPCError(err error, unknownMsg string) error {
	switch {
	case errors.Is(err, authsvc.ErrUserNotFound):
		return status.Errorf(codes.NotFound, "user not found")

	case errors.Is(err, authsvc.ErrUserWithEmailAlreadyExists):
		return status.Errorf(codes.AlreadyExists, "user with this email already exists")

	case errors.Is(err, authsvc.ErrUserWithLoginAlreadyExists):
		return status.Errorf(codes.AlreadyExists, "user with this login already exists")

	case errors.Is(err, authsvc.ErrUserEmailNotVerified):
		return status.Errorf(codes.PermissionDenied, "user has not verified email")

	case errors.Is(err, authsvc.ErrInvalidCredentials):
		return status.Errorf(codes.PermissionDenied, "invalid password")

	case errors.Is(err, authsvc.ErrTokenExpired):
		return status.Errorf(codes.PermissionDenied, "token is expired")

	case errors.Is(err, authsvc.ErrTokenMalformed):
		return status.Errorf(codes.PermissionDenied, "token is malformed")
	}

	return status.Errorf(codes.Internal, unknownMsg)
}
