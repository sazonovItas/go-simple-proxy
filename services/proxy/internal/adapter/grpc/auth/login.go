package grpcauth

import (
	"context"
	"fmt"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ar *authRepository) Login(ctx context.Context, login, password string) (string, error) {
	const op = "adapter.grpc.auth.Login"

	resp, err := ar.grpcAuthApi.Login(ctx, &authv1.LoginRequest{Login: login, Password: password})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.Token, nil
}
