package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
	accountsvc "github.com/sazonovItas/proxy-manager/services/auth/internal/service/account"
	accountv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/account/v1"
)

type AccountService interface {
	UserById(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type accountHandler struct {
	accountSvc AccountService

	accountv1.UnimplementedAccountServer
}

func Register(srv *grpc.Server, handler *accountHandler) {
	accountv1.RegisterAccountServer(srv, handler)
}

func New(accountSvc AccountService) *accountHandler {
	return &accountHandler{
		accountSvc: accountSvc,
	}
}

func (ah *accountHandler) User(
	ctx context.Context,
	r *accountv1.UserRequest,
) (*accountv1.UserResponse, error) {
	if r.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user id is required")
	}

	id, err := uuid.Parse(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user id bad format")
	}

	user, err := ah.accountSvc.UserById(ctx, id)
	if err != nil {
		if errors.Is(err, accountsvc.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to get user")
	}

	return &accountv1.UserResponse{
		Email:     user.Email,
		Login:     user.Login,
		Role:      user.UserRole,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}, nil
}
