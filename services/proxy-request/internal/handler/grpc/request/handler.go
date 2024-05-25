package grpcrequest

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

var ErrBadRequestUUID = errors.New("bad request uuid")

type requestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
	Request(ctx context.Context, id uuid.UUID) (*entity.Request, error)
	Timestamp(ctx context.Context, timestamp time.Time, limit int) ([]entity.Request, error)
}

type RequestHandler struct {
	l          *slog.Logger
	requestUsc requestUsecase

	requestv1.UnimplementedProxyRequestServiceServer
}

var _ requestv1.ProxyRequestServiceServer = (*RequestHandler)(nil)

func NewRequestHandler(logger *slog.Logger, requestUsc requestUsecase) *RequestHandler {
	return &RequestHandler{
		l:          logger,
		requestUsc: requestUsc,
	}
}
