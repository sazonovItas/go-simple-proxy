package service

import (
	"context"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
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
