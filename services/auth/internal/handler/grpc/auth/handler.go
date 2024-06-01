package auth

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
)

type AuthService interface {
	Register(ctx context.Context, email, login, password string) (uuid.UUID, error)
	Login(ctx context.Context, email, password string) (string, error)
	Validate(ctx context.Context, token string) (entity.UserInfo, error)

	VerifyEmail(ctx context.Context, verifyToken string) error
	ResetPassword(ctx context.Context, resetToken, password string) error
	GenerateResetToken(ctx context.Context, email string) error
	ValidateResetToken(ctx context.Context, resetToken string) error
}

type authHandler struct {
	auth AuthService

	authv1.UnimplementedAuthServer
}

func Register(srv *grpc.Server, handler *authHandler) {
	authv1.RegisterAuthServer(srv, handler)
}

func New(authSvc AuthService) *authHandler {
	return &authHandler{
		auth: authSvc,
	}
}
