package grpcapirequest

import (
	"context"

	"google.golang.org/grpc"

	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
)

type grpcRequestAPI interface {
	Save(
		ctx context.Context,
		in *requestv1.SaveRequest,
		opts ...grpc.CallOption,
	) (*requestv1.SaveResponse, error)
}

type RequestRepository struct {
	grpcRequestRepo grpcRequestAPI
}

func New(cli *grpc.ClientConn) *RequestRepository {
	return &RequestRepository{
		grpcRequestRepo: requestv1.NewProxyRequestServiceClient(cli),
	}
}
