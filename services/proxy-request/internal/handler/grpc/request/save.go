package grpcrequest

import (
	"context"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

// TODO: do not save requests with upload and download equal to 0
func (rh *RequestHandler) Save(
	ctx context.Context,
	in *requestv1.SaveRequest,
) (*requestv1.SaveResponse, error) {
	if in.Request.ProxyId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "proxy id is required")
	}

	if in.Request.ProxyUserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "proxy user id is required")
	}

	proxyRequest, err := ParseRequest(in.Request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err := rh.requestUsc.Save(ctx, proxyRequest); err != nil {
		rh.l.Error("failed save proxy request", slogger.Err(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &requestv1.SaveResponse{Id: proxyRequest.ID.String()}, nil
}
