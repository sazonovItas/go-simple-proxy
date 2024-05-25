package grpcrequest

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

func (ph *RequestHandler) ProxyRequestByTimestamp(
	ctx context.Context,
	in *requestv1.TimestampRequest,
) (*requestv1.TimestampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Timestamp is unimplemented")
}
