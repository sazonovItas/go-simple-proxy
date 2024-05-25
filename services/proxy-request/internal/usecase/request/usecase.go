package requestusc

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

type requestRepository interface {
	Save(ctx context.Context, request *entity.Request) error
	Request(ctx context.Context, id uuid.UUID) (*entity.Request, error)
	Timestamp(ctx context.Context, timestamp time.Time, limit int) ([]entity.Request, error)

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
