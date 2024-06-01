package proxysvc

import (
	"context"
	"log/slog"
	"time"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

type requestRepository interface {
	Save(ctx context.Context, request *entity.Request) error
}

type authRepository interface {
	Login(ctx context.Context, login, password string) (string, error)
	Validate(ctx context.Context, token string) (string, error)
}

type tokenRepository[T any] interface {
	Get(key string) (T, error)
	Set(key string, value T, duration time.Duration)
	Delete(keys ...string)
}

type ProxyService struct {
	log *slog.Logger

	authRepo    authRepository
	requestRepo requestRepository
	tokenRepo   tokenRepository[entity.Token]
}

func New(
	l *slog.Logger,

	authRepo authRepository,
	requestRepo requestRepository,
	tokenRepo tokenRepository[entity.Token],
) *ProxyService {
	return &ProxyService{
		log: l,

		authRepo:    authRepo,
		tokenRepo:   tokenRepo,
		requestRepo: requestRepo,
	}
}
