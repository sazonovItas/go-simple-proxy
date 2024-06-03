package grpcuserapi

import (
	"context"

	accountv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/account/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

func (ua *userApi) User(ctx context.Context, id string) (*entity.User, error) {
	const op = "adapter.grpcapi.user.User"

	u, err := ua.accountApi.User(ctx, &accountv1.UserRequest{Id: id})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return nil, adapter.ErrInvalidArgument
			case codes.NotFound:
				return nil, adapter.ErrUserNotFound
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &entity.User{
		Email:     u.GetEmail(),
		Login:     u.GetLogin(),
		Role:      u.GetRole(),
		CreatedAt: u.GetCreatedAt().AsTime(),
	}, nil
}
