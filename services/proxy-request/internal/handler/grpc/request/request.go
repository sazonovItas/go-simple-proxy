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

func (rh *requestHandler) Request(
	ctx context.Context,
	in *requestv1.GetRequest,
) (*requestv1.GetResponse, error) {
	if in.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "proxy request id is required")
	}

	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad proxy request id format")
	}

	r, err := rh.requestUsc.Request(ctx, id)
	if err != nil {
		if errors.Is(err, adapter.ErrRequestNotFound) {
			return nil, status.Errorf(codes.NotFound, "no proxy request is found")
		}

		return nil, status.Errorf(codes.Internal, "failed to get proxy request")
	}

	return &requestv1.GetResponse{Request: ParseProxyRequest(r)}, nil
}
