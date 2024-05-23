package grpcrequest

import (
	"context"

	"github.com/google/uuid"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

func (rh *RequestHandler) Save(
	ctx context.Context,
	r *requestv1.SaveRequest,
) (*requestv1.SaveResponse, error) {
	proxyId, _ := uuid.Parse(r.Request.ProxyId)
	proxyUserId, _ := uuid.Parse(r.Request.ProxyUserName)

	proxyRequest := entity.Request{
		ProxyID:       proxyId,
		ProxyName:     r.Request.ProxyName,
		ProxyUserID:   proxyUserId,
		ProxyUserIP:   r.Request.ProxyUserIp,
		ProxyUserName: r.Request.ProxyUserName,
		Host:          r.Request.Host,
		Upload:        r.Request.Upload,
		Download:      r.Request.Download,
		CreatedAt:     r.Request.CreatedAt.AsTime(),
	}
	rh.l.Info("new request", "request", proxyRequest)

	if err := rh.requestUsc.Save(ctx, &proxyRequest); err != nil {
		rh.l.Error("failed save proxy request", slogger.Err(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &requestv1.SaveResponse{Id: proxyRequest.ID.String()}, nil
}
