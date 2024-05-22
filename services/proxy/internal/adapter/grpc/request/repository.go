package grpcrequest

import (
	request "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb"
	"google.golang.org/grpc"
)

type rpcRequestService interface{}

type RequestRepository struct {
	cli request.ProxyRequestServiceClient
}

func NewRequestRepository(cli *grpc.ClientConn) *RequestRepository {
	return &RequestRepository{
		cli: request.NewProxyRequestServiceClient(cli),
	}
}
