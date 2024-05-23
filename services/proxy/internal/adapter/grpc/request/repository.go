package requestrepo

import (
	"context"

	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
	"google.golang.org/grpc"
)

type rpcRequestRepository interface {
	Save(
		ctx context.Context,
		in *requestv1.SaveRequest,
		opts ...grpc.CallOption,
	) (*requestv1.SaveResponse, error)
}

type RequestRepository struct {
	rpcRequestRepo rpcRequestRepository
}

func NewRequestRepository(cli *grpc.ClientConn) *RequestRepository {
	return &RequestRepository{
		rpcRequestRepo: requestv1.NewProxyRequestServiceClient(cli),
	}
}
