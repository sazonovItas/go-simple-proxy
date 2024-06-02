package grpcuserapi

import (
	"context"
	"fmt"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
)

func (ua *userApi) Register(
	ctx context.Context,
	email, login, password string,
) (string, error) {
	const op = "adapter.grpcapi.user.Register"

	resp, err := ua.authApi.Register(
		ctx,
		&authv1.RegisterRequest{Email: email, Login: login, Password: password},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return "", fmt.Errorf("%s: %w", op, adapter.ErrInvalidArgument)
			case codes.AlreadyExists:
				return "", fmt.Errorf("%s: %w", op, adapter.ErrUserAlreadyExists)
			default:
				return "", fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return "", fmt.Errorf("%s: %w", op, err)
		}
	}

	return resp.GetId(), nil
}
