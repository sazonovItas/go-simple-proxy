package grpcrequest

import (
	"context"
	"fmt"

	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

func (rr *RequestRepository) Save(ctx context.Context, r *entity.Request) error {
	const op = "adapter.grpc.grpcrequest.Save"

	saveRequest := &requestv1.SaveRequest{
		Request: &requestv1.ProxyRequest{
			ProxyId:       r.ProxyID,
			ProxyName:     r.ProxyName,
			ProxyUserId:   r.ProxyUserID,
			ProxyUserIp:   r.ProxyUserIP,
			ProxyUserName: r.ProxyUserName,
			Host:          r.Host,
			Upload:        r.Upload,
			Download:      r.Download,
			CreatedAt:     timestamppb.New(r.CreatedAt),
		},
	}

	resp, err := rr.grpcRequestRepo.Save(ctx, saveRequest)
	if err != nil {
		return fmt.Errorf("%s: failed save proxy request: %w", op, err)
	}
	r.ID = resp.GetId()

	return nil
}
