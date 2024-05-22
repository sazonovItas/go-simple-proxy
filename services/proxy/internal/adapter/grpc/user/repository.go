package grpcuser

import "google.golang.org/grpc"

type RequestRepository struct {
	cli *grpc.ClientConn
}

func NewRequestRepository(cli *grpc.ClientConn) *RequestRepository {
	return &RequestRepository{
		cli: cli,
	}
}
