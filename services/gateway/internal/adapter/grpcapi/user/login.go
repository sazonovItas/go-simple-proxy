package grpcuserapi

import (
	"context"
	"fmt"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
)

func (ua *userApi) Login(
	ctx context.Context,
	login, password string,
) (string, error) {
	const op = "adapter.grpcapi.user.Login"

	resp, err := ua.authApi.Login(
		ctx,
		&authv1.LoginRequest{Login: login, Password: password},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return "", fmt.Errorf("%s: %w", op, adapter.ErrInvalidArgument)
			case codes.PermissionDenied:
				return "", fmt.Errorf("%s: %w", op, adapter.ErrUserInvalidaCreditanals)
			case codes.NotFound:
				return "", fmt.Errorf("%s: %w", op, adapter.ErrUserNotFound)
			default:
				return "", fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return "", fmt.Errorf("%s: %w", op, err)
		}
	}

	return resp.GetToken(), nil
}
