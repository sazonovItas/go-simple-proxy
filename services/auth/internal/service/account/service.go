package accountsvc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

var ErrUserNotFound = errors.New("user not found")

type userRepo interface {
	UserById(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type accountService struct {
	userRepo userRepo

	log *slog.Logger
}

func New(userRepo userRepo, l *slog.Logger) *accountService {
	return &accountService{
		userRepo: userRepo,
		log:      l,
	}
}

func (as *accountService) UserById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	const op = "service.account.UserById"

	as.log.Info("attemting get user by id")

	user, err := as.userRepo.UserById(ctx, id)
	if err != nil {
		as.log.Error("failed get user by id", slogger.Err(err))

		if errors.Is(err, adapter.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
