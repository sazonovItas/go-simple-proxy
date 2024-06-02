package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

func (ah *authHandler) Register(
	ctx context.Context,
	r *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {
	if r.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email is required")
	}

	if r.GetLogin() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login is required")
	}

	if r.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is required")
	}

	id, err := ah.auth.Register(ctx, r.GetEmail(), r.GetLogin(), r.GetPassword())
	if err != nil {
		return nil, GRPCError(err, "failed to register")
	}

	return &authv1.RegisterResponse{Id: id.String()}, nil
}
