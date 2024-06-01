package authsvc

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

type userRepo interface {
	NewUser(ctx context.Context, user *entity.User) (uuid.UUID, error)
	UserByEmail(ctx context.Context, email string) (*entity.User, error)
	UserByLogin(ctx context.Context, login string) (*entity.User, error)
	UserByVerifyToken(ctx context.Context, verifyToken string) (*entity.User, error)
	UserByResetToken(ctx context.Context, resetToken string) (*entity.User, error)

	NewResetToken(ctx context.Context, email, resetToken string) error
	VerifyEmail(ctx context.Context, verifyToken string) error
	UpdatePasswordByResetToken(ctx context.Context, resetToken, passwordHash string) error
}

type Hasher interface {
	PasswordHash(password string) ([]byte, error)
	Compare(hashedPassword string, password string) error
}

type authService struct {
	userRepo userRepo
	hasher   Hasher

	log        *slog.Logger
	authSecret string
	tokenTTL   time.Duration
}

func New(
	userRepo userRepo,
	hasher Hasher,
	l *slog.Logger,
	authSecret string,
	tokenTTL time.Duration,
) *authService {
	return &authService{
		log:        l,
		authSecret: authSecret,
		tokenTTL:   tokenTTL,

		userRepo: userRepo,
		hasher:   hasher,
	}
}
