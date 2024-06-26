package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ah *authHandler) Login(
	ctx context.Context,
	r *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {
	if r.GetLogin() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login is required")
	}

	if r.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is required")
	}

	token, err := ah.auth.Login(ctx, r.GetLogin(), r.GetPassword())
	if err != nil {
		return nil, GRPCError(err, "failed to login")
	}

	return &authv1.LoginResponse{Token: token}, nil
}
