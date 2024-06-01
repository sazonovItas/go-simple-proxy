package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ah *authHandler) Validate(
	ctx context.Context,
	r *authv1.ValidateRequest,
) (*authv1.ValidateResponse, error) {
	if r.Token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "token is required")
	}

	claims, err := ah.auth.Validate(ctx, r.Token)
	if err != nil {
		return nil, GRPCError(err, "failed validate token")
	}

	return &authv1.ValidateResponse{
		Id:    claims.ID,
		Email: claims.Email,
		Login: claims.Login,
		Role:  claims.Role,
	}, nil
}
