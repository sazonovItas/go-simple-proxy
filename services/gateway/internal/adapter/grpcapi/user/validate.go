package grpcuserapi

import (
	"context"
	"fmt"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

func (ua *userApi) ValidateToken(ctx context.Context, token string) (*entity.Token, error) {
	const op = "adapter.grpcapi.user.ValidateToken"

	resp, err := ua.authApi.Validate(ctx, &authv1.ValidateRequest{Token: token})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrInvalidArgument)
			case codes.AlreadyExists:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrUserAlreadyExists)
			default:
				return nil, fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return &entity.Token{
		ID: resp.GetId(),
	}, nil
}
