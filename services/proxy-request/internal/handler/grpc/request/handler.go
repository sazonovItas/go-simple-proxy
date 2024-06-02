package grpcrequest

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/entity"
	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
)

var ErrBadRequestUUID = errors.New("bad request uuid")

type RequestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
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

type requestHandler struct {
	l          *slog.Logger
	requestUsc RequestUsecase

	requestv1.UnimplementedProxyRequestServiceServer
}

var _ requestv1.ProxyRequestServiceServer = (*requestHandler)(nil)

func Register(srv *grpc.Server, handler *requestHandler) {
	requestv1.RegisterProxyRequestServiceServer(srv, handler)
}

func New(logger *slog.Logger, requestUsc RequestUsecase) *requestHandler {
	return &requestHandler{
		l:          logger,
		requestUsc: requestUsc,
	}
}
