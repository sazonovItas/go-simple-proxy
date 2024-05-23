package grpcrequest

import (
	"context"
	"log/slog"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

type requestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
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
