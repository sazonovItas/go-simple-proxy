package grpcuserapi

import (
	"context"

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
		&authv1.LoginRequest{
			Login:    login,
			Password: password,
		},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return "", adapter.ErrInvalidArgument
			case codes.PermissionDenied:
				return "", adapter.ErrUserInvalidCreditanals
			case codes.NotFound:
				return "", adapter.ErrUserNotFound
			default:
				return "", err
			}
		} else {
			return "", err
		}
	}

	return resp.GetToken(), nil
}
