package grpcuserapi

import (
	"context"

	accountv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/account/v1"
	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
	"google.golang.org/grpc"
)

type grpcAuthApi interface {
	Register(
		ctx context.Context,
		in *authv1.RegisterRequest,
		opts ...grpc.CallOption,
	) (*authv1.RegisterResponse, error)

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

type grpcAccountApi interface {
	User(
		ctx context.Context,
		in *accountv1.UserRequest,
		opts ...grpc.CallOption,
	) (*accountv1.UserResponse, error)
}

type userApi struct {
	accountApi grpcAccountApi
	authApi    grpcAuthApi
}

func New(authApi grpcAuthApi, accountApi grpcAccountApi) *userApi {
	return &userApi{
		authApi:    authApi,
		accountApi: accountApi,
	}
}
