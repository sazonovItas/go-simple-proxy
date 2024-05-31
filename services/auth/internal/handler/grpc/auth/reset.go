package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ah *authHandler) ResetToken(
	ctx context.Context,
	r *authv1.ResetPasswordRequest,
) (*authv1.Empty, error) {
	if r.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email is required")
	}

	err := ah.auth.ResetToken(ctx, r.Email)
	if err != nil {
		return nil, GRPCError(err, "failed generate reset token")
	}

	return &authv1.Empty{}, nil
}

func (ah *authHandler) UpdatePassword(
	ctx context.Context,
	r *authv1.UpdateResetPasswordRequest,
) (*authv1.Empty, error) {
	if r.Token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "reset token is required")
	}

	if r.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is required")
	}

	err := ah.auth.UpdatePassword(ctx, r.Password, r.Token)
	if err != nil {
		return nil, GRPCError(err, "failed update password")
	}

	return &authv1.Empty{}, nil
}
