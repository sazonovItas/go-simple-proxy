package grpcrequest

import (
	"context"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
	pb_request "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb"
)

func (rh *RequestHandler) SaveProxyRequest(
	ctx context.Context,
	r *pb_request.SaveRequest,
) (*pb_request.SaveResponse, error) {
	proxyRequest := entity.Request{
		ProxyID:       r.Request.ProxyId,
		ProxyName:     r.Request.ProxyName,
		ProxyUserID:   r.Request.ProxyUserId,
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

	return &pb_request.SaveResponse{Id: proxyRequest.ID}, nil
}
