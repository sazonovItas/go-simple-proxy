package proxysvc

import (
	"context"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

type requestRepository interface {
	Save(ctx context.Context, request *entity.Request) error
}

type userRepository interface {
	UserByName(ctx context.Context, name string) (*entity.User, error)
}

type ProxyService struct {
	requestRepo requestRepository
	userRepo    userRepository
}

func New(requestRepo requestRepository, userRepo userRepository) *ProxyService {
	return &ProxyService{
		userRepo:    userRepo,
		requestRepo: requestRepo,
	}
}
