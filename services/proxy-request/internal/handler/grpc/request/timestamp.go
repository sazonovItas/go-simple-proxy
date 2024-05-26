package grpcrequest

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

func (ph *requestHandler) ProxyRequestByTimestamp(
	ctx context.Context,
	in *requestv1.TimestampRequest,
) (*requestv1.TimestampResponse, error) {
	if in.Limit <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "limit should be greater than 0")
	}

	if in.Timestamp == nil {
		return nil, status.Errorf(codes.InvalidArgument, "timestamp is required")
	}

	requests, err := ph.requestUsc.Timestamp(ctx, in.Timestamp.AsTime(), int(in.Limit))
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return nil, status.Errorf(codes.NotFound, "no requests found")
		default:
			return nil, status.Errorf(codes.Internal, "internal error")
		}
	}

	proxyRequests := make([]*requestv1.ProxyRequest, 0, len(requests))
	for _, r := range requests {
		proxyRequests = append(proxyRequests, ParseProxyRequest(&r))
	}

	return &requestv1.TimestampResponse{Requests: proxyRequests}, nil
}
