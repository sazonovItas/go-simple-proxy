package grpcapirequest

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

func (rr *RequestRepository) Save(ctx context.Context, r *entity.Request) error {
	const op = "adapter.grpc.grpcrequest.Save"

	saveRequest := &requestv1.SaveRequest{
		Request: &requestv1.ProxyRequest{
			UserId:    r.UserID,
			ProxyId:   r.ProxyID,
			RemoteIp:  r.RemoteIP,
			Host:      r.Host,
			Upload:    r.Upload,
			Download:  r.Download,
			CreatedAt: timestamppb.New(r.CreatedAt),
		},
	}

	resp, err := rr.grpcRequestRepo.Save(ctx, saveRequest)
	if err != nil {
		return fmt.Errorf("%s: failed save proxy request: %w", op, err)
	}
	r.ID = resp.GetId()

	return nil
}
