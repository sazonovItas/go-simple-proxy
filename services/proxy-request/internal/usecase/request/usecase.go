package requestusc

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/entity"
)

var ErrRequestNotFound = errors.New("request not found")

type requestRepository interface {
	Save(ctx context.Context, request *entity.Request) error
	Request(ctx context.Context, id uuid.UUID) (*entity.Request, error)
	Timestamp(ctx context.Context, from, to time.Time) ([]entity.Request, error)

	TimestampAndUserId(
		ctx context.Context,
		from, to time.Time,
		userId uuid.UUID,
	) ([]entity.Request, error)

	TimestampAndProxyId(
		ctx context.Context,
		from, to time.Time,
		proxyId uuid.UUID,
	) ([]entity.Request, error)
}

type requestUsecase struct {
	requestRepo requestRepository
}

func New(requestRepo requestRepository) *requestUsecase {
	return &requestUsecase{
		requestRepo: requestRepo,
	}
}

func (ru *requestUsecase) Save(ctx context.Context, request *entity.Request) error {
	return ru.requestRepo.Save(ctx, request)
}

func (ru *requestUsecase) Request(ctx context.Context, id uuid.UUID) (*entity.Request, error) {
	r, err := ru.requestRepo.Request(ctx, id)
	if err != nil && errors.Is(err, adapter.ErrRequestNotFound) {
		err = ErrRequestNotFound
	}

	return r, err
}

func (ru *requestUsecase) Timestamp(
	ctx context.Context,
	from, to time.Time,
) ([]entity.Request, error) {
	rs, err := ru.requestRepo.Timestamp(ctx, from, to)
	if err != nil && errors.Is(err, adapter.ErrRequestNotFound) {
		err = ErrRequestNotFound
	}

	return rs, err
}

func (ru *requestUsecase) TimestampAndUserId(
	ctx context.Context,
	from, to time.Time,
	userId uuid.UUID,
) ([]entity.Request, error) {
	rs, err := ru.requestRepo.TimestampAndUserId(ctx, from, to, userId)
	if err != nil && errors.Is(err, adapter.ErrRequestNotFound) {
		err = ErrRequestNotFound
	}

	return rs, err
}

func (ru *requestUsecase) TimestampAndProxyId(
	ctx context.Context,
	from, to time.Time,
	proxyId uuid.UUID,
) ([]entity.Request, error) {
	rs, err := ru.requestRepo.TimestampAndProxyId(ctx, from, to, proxyId)
	if err != nil && errors.Is(err, adapter.ErrRequestNotFound) {
		err = ErrRequestNotFound
	}

	return rs, err
}
