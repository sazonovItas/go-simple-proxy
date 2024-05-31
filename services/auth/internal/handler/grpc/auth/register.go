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
	if r.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email is required")
	}

	if r.Login == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login is required")
	}

	if r.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is required")
	}

	id, err := ah.auth.Register(ctx, r.Email, r.Login, r.Password)
	if err != nil {
		return nil, GRPCError(err, "failed to register")
	}

	return &authv1.RegisterResponse{Id: id.String()}, nil
}
