package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ah *authHandler) VerifyEmail(
	ctx context.Context,
	r *authv1.VerifyEmailRequest,
) (*authv1.Empty, error) {
	if r.GetVerifyToken() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "verify token is required")
	}

	err := ah.auth.VerifyEmail(ctx, r.GetVerifyToken())
	if err != nil {
		return nil, GRPCError(err, "failed to verify email")
	}

	return &authv1.Empty{}, nil
}
