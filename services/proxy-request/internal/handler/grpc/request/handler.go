package grpcrequest

import (
	"context"
	"log/slog"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
	pb_request "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb"
)

type requestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
}

type RequestHandler struct {
	l          *slog.Logger
	requestUsc requestUsecase

	pb_request.UnimplementedProxyRequestServiceServer
}

var _ pb_request.ProxyRequestServiceServer = (*RequestHandler)(nil)

func NewRequestHandler(logger *slog.Logger, requestUsc requestUsecase) *RequestHandler {
	return &RequestHandler{
		l:          logger,
		requestUsc: requestUsc,
	}
}
