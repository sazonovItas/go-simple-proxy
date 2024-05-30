package authsvc

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/entity"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotVerified    = errors.New("user has not verified email")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type userRepo interface {
	Create(ctx context.Context, user *entity.User) (uuid.UUID, error)
	UserByEmail(ctx context.Context, email string) (*entity.User, error)
	UserByLogin(ctx context.Context, login string) (*entity.User, error)
	UserExists(ctx context.Context, email, login string) (*entity.User, error)
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
