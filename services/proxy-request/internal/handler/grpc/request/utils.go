package grpcrequest

import (
	"errors"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/entity"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
)

var (
	ErrBadProxyIdFormat     = errors.New("bad proxy id format")
	ErrBadProxyUserIdFormat = errors.New("bad proxy user id format")
)

func ParseRequest(r *requestv1.ProxyRequest) (*entity.Request, error) {
	proxyId, err := uuid.Parse(r.ProxyId)
	if err != nil {
		return nil, ErrBadProxyIdFormat
	}

	proxyUserId, err := uuid.Parse(r.ProxyUserId)
	if err != nil {
		return nil, ErrBadProxyUserIdFormat
	}

	return &entity.Request{
		ProxyID:       proxyId,
		ProxyName:     r.ProxyName,
		ProxyUserID:   proxyUserId,
		ProxyUserIP:   r.ProxyUserIp,
		ProxyUserName: r.ProxyUserName,
		Host:          r.Host,
		Upload:        r.Upload,
		Download:      r.Download,
		CreatedAt:     r.CreatedAt.AsTime(),
	}, nil
}

func ParseProxyRequest(r *entity.Request) *requestv1.ProxyRequest {
	return &requestv1.ProxyRequest{
		Id:            r.ID.String(),
		ProxyId:       r.ProxyID.String(),
		ProxyName:     r.ProxyName,
		ProxyUserId:   r.ProxyUserID.String(),
		ProxyUserIp:   r.ProxyUserIP,
		ProxyUserName: r.ProxyUserName,
		Host:          r.Host,
		Upload:        r.Upload,
		Download:      r.Download,
		CreatedAt:     timestamppb.New(r.CreatedAt),
	}
}
