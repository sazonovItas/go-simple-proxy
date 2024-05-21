package request

import (
	"context"
	"time"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

type requestRepository interface {
	Save(ctx context.Context, request *entity.Request) error
	GetByID(ctx context.Context, id string) (entity.Request, error)
	GetByTimestamp(ctx context.Context, timestamp time.Time, limit int) ([]entity.Request, error)

	GetByHostAndTimestamp(
		ctx context.Context,
		timestamp time.Time,
		host string,
		limit int,
	) ([]entity.Request, error)

	GetByProxyUserIDAndTimestamp(
		ctx context.Context,
		timestamp time.Time,
		proxyId string,
		limit int,
	) ([]entity.Request, error)

	GetByProxyIDAndTimestamp(
		ctx context.Context,
		timestamp time.Time,
		proxyId string,
		limit int,
	) ([]entity.Request, error)
}

type RequestUsecase struct {
	requestRepo requestRepository
}

func NewRequestUsecase(requestRepo requestRepository) *RequestUsecase {
	return &RequestUsecase{
		requestRepo: requestRepo,
	}
}
