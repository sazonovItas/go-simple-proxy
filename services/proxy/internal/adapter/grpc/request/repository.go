package grpcrequest

import (
	"context"

	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
	"google.golang.org/grpc"
)

type grpcRequestRepository interface {
	Save(
		ctx context.Context,
		in *requestv1.SaveRequest,
		opts ...grpc.CallOption,
	) (*requestv1.SaveResponse, error)
}

type RequestRepository struct {
	grpcRequestRepo grpcRequestRepository
}

func New(grpcRequestRepo grpcRequestRepository) *RequestRepository {
	return &RequestRepository{
		grpcRequestRepo: grpcRequestRepo,
	}
}
