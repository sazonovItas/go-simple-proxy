package grpcrequest

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

func (ph *requestHandler) Timestamp(
	ctx context.Context,
	in *requestv1.TimestampRequest,
) (*requestv1.TimestampResponse, error) {
	if in.From == nil {
		return nil, status.Errorf(codes.InvalidArgument, "from timestamp is required")
	}

	if in.To == nil {
		return nil, status.Errorf(codes.InvalidArgument, "to timestamp is required")
	}

	requests, err := ph.requestUsc.Timestamp(ctx, in.From.AsTime(), in.To.AsTime())
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return nil, status.Errorf(codes.NotFound, "no requests found")
		default:
			return nil, status.Errorf(codes.Internal, "internal error")
		}
	}

	proxyRequests := make([]*requestv1.ProxyRequest, 0, len(requests))
	for i := range requests {
		proxyRequests = append(proxyRequests, ParseProxyRequest(&requests[i]))
	}

	return &requestv1.TimestampResponse{Requests: proxyRequests}, nil
}
