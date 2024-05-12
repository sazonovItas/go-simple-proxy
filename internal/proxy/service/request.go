package service

import (
	"context"

	"github.com/sazonovItas/go-simple-proxy/internal/proxy/entity"
)

type requestRepository interface {
	Save(ctx context.Context, request entity.Request) error
}

type RequestService struct {
	requestRepo requestRepository
}

func NewRequestService(requestRepository requestRepository) *RequestService {
	return &RequestService{
		requestRepo: requestRepository,
	}
}
