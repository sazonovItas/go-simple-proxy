package grpcrequestapi

import (
	"context"

	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
	"google.golang.org/grpc"
)

type grpcRequestApi interface {
	Timestamp(
		ctx context.Context,
		in *requestv1.TimestampRequest,
		opts ...grpc.CallOption,
	) (*requestv1.TimestampResponse, error)

	TimestampAndUserId(
		ctx context.Context,
		in *requestv1.TimestampAndIdRequest,
		opts ...grpc.CallOption,
	) (*requestv1.TimestampResponse, error)

	TimestampAndProxyId(
		ctx context.Context,
		in *requestv1.TimestampAndIdRequest,
		opts ...grpc.CallOption,
	) (*requestv1.TimestampResponse, error)
}

type requestApi struct {
	reqApi grpcRequestApi
}

func New(reqApi grpcRequestApi) *requestApi {
	return &requestApi{
		reqApi: reqApi,
	}
}
