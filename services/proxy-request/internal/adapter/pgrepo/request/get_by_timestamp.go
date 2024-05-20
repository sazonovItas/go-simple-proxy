package request

import (
	"context"
	"time"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
)

func (rr *RequestRepository) GetByTimestamp(
	ctx context.Context,
	timestamp time.Time,
	limit int,
) ([]entity.Request, error) {
	return nil, nil
}

func (rr *RequestRepository) GetByProxyUser(
	ctx context.Context,
	ProxyUser string,
	timestamp time.Time,
	limit int,
) ([]entity.Request, error) {
	return nil, nil
}

func (rr *RequestRepository) GetByProxy(
	ctx context.Context,
	Proxy string,
	timestamp time.Time,
	limit int,
) ([]entity.Request, error) {
	return nil, nil
}

func (rr *RequestRepository) GetByHost(
	ctx context.Context,
	host string,
	timestamp time.Time,
	limit int,
) ([]entity.Request, error) {
	return nil, nil
}
