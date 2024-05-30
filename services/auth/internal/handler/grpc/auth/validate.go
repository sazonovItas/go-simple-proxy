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
) (*authv1.Empty, error) {
	if r.Token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "token is required")
	}

	if err := ah.auth.Validate(ctx, r.Token); err != nil {
		return nil, GRPCError(err)
	}

	return &authv1.Empty{}, nil
}
