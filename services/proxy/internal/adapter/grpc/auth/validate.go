package grpcauth

import (
	"context"
	"fmt"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ar *authRepository) Validate(ctx context.Context, token string) error {
	const op = "adapter.grpc.auth.Validate"

	_, err := ar.grpcAuthApi.Validate(ctx, &authv1.ValidateRequest{Token: token})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
