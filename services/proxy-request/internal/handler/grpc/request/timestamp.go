package grpcrequest

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/adapter"
	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
)

func (ph *requestHandler) Timestamp(
	ctx context.Context,
	in *requestv1.TimestampRequest,
) (*requestv1.TimestampResponse, error) {
	if in.GetFrom() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "from timestamp is required")
	}

	if in.GetTo() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "to timestamp is required")
	}

	requests, err := ph.requestUsc.Timestamp(ctx, in.GetFrom().AsTime(), in.GetTo().AsTime())
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return nil, status.Errorf(codes.NotFound, "no requests found")
		default:
			return nil, status.Errorf(codes.Internal, "failed to get requests")
		}
	}

	proxyRequests := make([]*requestv1.ProxyRequest, 0, len(requests))
	for i := range requests {
		proxyRequests = append(proxyRequests, ParseProxyRequest(&requests[i]))
	}

	return &requestv1.TimestampResponse{Requests: proxyRequests}, nil
}

func (ph *requestHandler) TimestampAndUserId(
	ctx context.Context,
	in *requestv1.TimestampAndIdRequest,
) (*requestv1.TimestampResponse, error) {
	if in.GetFrom() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "from timestamp is required")
	}

	if in.GetTo() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "to timestamp is required")
	}

	if in.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user id is required")
	}

	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user id bad format")
	}

	requests, err := ph.requestUsc.TimestampAndUserId(
		ctx,
		in.GetFrom().AsTime(),
		in.GetTo().AsTime(),
		id,
	)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return nil, status.Errorf(codes.NotFound, "no requests found")
		default:
			return nil, status.Errorf(codes.Internal, "failed to get requests")
		}
	}

	proxyRequests := make([]*requestv1.ProxyRequest, 0, len(requests))
	for i := range requests {
		proxyRequests = append(proxyRequests, ParseProxyRequest(&requests[i]))
	}

	return &requestv1.TimestampResponse{Requests: proxyRequests}, nil
}

func (ph *requestHandler) TimestampAndProxyId(
	ctx context.Context,
	in *requestv1.TimestampAndIdRequest,
) (*requestv1.TimestampResponse, error) {
	if in.GetFrom() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "from timestamp is required")
	}

	if in.GetTo() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "to timestamp is required")
	}

	if in.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "proxy id is required")
	}

	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user id bad format")
	}

	requests, err := ph.requestUsc.TimestampAndProxyId(
		ctx,
		in.GetFrom().AsTime(),
		in.GetTo().AsTime(),
		id,
	)
	if err != nil {
		switch {
		case errors.Is(err, adapter.ErrRequestNotFound):
			return nil, status.Errorf(codes.NotFound, "no requests found")
		default:
			return nil, status.Errorf(codes.Internal, "failed to get requests")
		}
	}

	proxyRequests := make([]*requestv1.ProxyRequest, 0, len(requests))
	for i := range requests {
		proxyRequests = append(proxyRequests, ParseProxyRequest(&requests[i]))
	}

	return &requestv1.TimestampResponse{Requests: proxyRequests}, nil
}
