package grpcrequest

import (
	"context"

	pb_request "github.com/sazonovItas/proxy-manager/proxy-request/api/proto/pb"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

type requestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
}

type RequestHandler struct {
	RequestUsc requestUsecase

	pb_request.UnimplementedProxyRequestServiceServer
}

var _ pb_request.ProxyRequestServiceServer = (*RequestHandler)(nil)

func NewRequestHandler(requestUsc requestUsecase) *RequestHandler {
	return &RequestHandler{
		RequestUsc: requestUsc,
	}
}
