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

	userId, err := uuid.Parse(r.UserId)
	if err != nil {
		return nil, ErrBadProxyUserIdFormat
	}

	return &entity.Request{
		UserID:    userId,
		ProxyID:   proxyId,
		RemoteIP:  r.RemoteIp,
		Host:      r.Host,
		Upload:    r.Upload,
		Download:  r.Download,
		CreatedAt: r.CreatedAt.AsTime(),
	}, nil
}

func ParseProxyRequest(r *entity.Request) *requestv1.ProxyRequest {
	return &requestv1.ProxyRequest{
		Id:        r.ID.String(),
		UserId:    r.UserID.String(),
		ProxyId:   r.ProxyID.String(),
		RemoteIp:  r.RemoteIP,
		Host:      r.Host,
		Upload:    r.Upload,
		Download:  r.Download,
		CreatedAt: timestamppb.New(r.CreatedAt),
	}
}
