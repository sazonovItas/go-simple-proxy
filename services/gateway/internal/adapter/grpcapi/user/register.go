package grpcuserapi

import (
	"context"

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
				return "", adapter.ErrInvalidArgument
			case codes.AlreadyExists:
				return "", adapter.ErrUserAlreadyExists
			default:
				return "", err
			}
		} else {
			return "", err
		}
	}

	return resp.GetId(), nil
}
