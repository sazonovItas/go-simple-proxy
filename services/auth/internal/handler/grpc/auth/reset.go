package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ah *authHandler) GenerateResetToken(
	ctx context.Context,
	r *authv1.GenerateResetTokenRequest,
) (*authv1.Empty, error) {
	if r.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email is required")
	}

	err := ah.auth.GenerateResetToken(ctx, r.GetEmail())
	if err != nil {
		return nil, GRPCError(err, "failed to generate reset token")
	}

	return &authv1.Empty{}, nil
}

func (ah *authHandler) ValidateResetToken(
	ctx context.Context,
	r *authv1.ValidateResetTokenRequest,
) (*authv1.Empty, error) {
	if r.GetResetToken() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "reset token is required")
	}

	err := ah.auth.ValidateResetToken(ctx, r.GetResetToken())
	if err != nil {
		return nil, GRPCError(err, "failed to validate reset token")
	}

	return &authv1.Empty{}, nil
}

func (ah *authHandler) ResetPassword(
	ctx context.Context,
	r *authv1.ResetPasswordRequest,
) (*authv1.Empty, error) {
	if r.GetResetToken() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "reset token is required")
	}

	if r.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is required")
	}

	err := ah.auth.ResetPassword(ctx, r.GetResetToken(), r.GetPassword())
	if err != nil {
		return nil, GRPCError(err, "failed to reset password")
	}

	return &authv1.Empty{}, nil
}
