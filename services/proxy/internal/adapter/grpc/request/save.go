package grpcrequest

import (
	"context"
	"fmt"

	request "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

func (rr *RequestRepository) Save(ctx context.Context, r *entity.Request) error {
	const op = "internal.adapter.grpc.grpcrequest.Save"

	saveRequest := &request.SaveRequest{
		Request: &request.ProxyRequest{
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

	resp, err := rr.cli.SaveProxyRequest(ctx, saveRequest)
	if err != nil {
		return fmt.Errorf("%s: failed save proxy request: %w", op, err)
	}
	r.ID = resp.Id

	return nil
}
