package grpcauth

import (
	"context"

	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
	"google.golang.org/grpc"
)

type grpcAuthAPI interface {
	Login(
		ctx context.Context,
		in *authv1.LoginRequest,
		opts ...grpc.CallOption,
	) (*authv1.LoginResponse, error)

	Validate(
		ctx context.Context,
		in *authv1.ValidateRequest,
		opts ...grpc.CallOption,
	) (*authv1.ValidateResponse, error)
}

type authRepository struct {
	grpcAuthApi grpcAuthAPI
}

func New(cli *grpc.ClientConn) *authRepository {
	return &authRepository{
		grpcAuthApi: authv1.NewAuthClient(cli),
	}
}
