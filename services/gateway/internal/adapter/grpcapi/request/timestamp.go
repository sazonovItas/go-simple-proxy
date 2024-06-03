package grpcrequestapi

import (
	"context"
	"fmt"
	"time"

	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

func (ra *requestApi) Timestamp(
	ctx context.Context,
	from, to time.Time,
) ([]*entity.Request, error) {
	const op = "adapter.grpcapi.user.Timestamp"

	requests, err := ra.reqApi.Timestamp(
		ctx,
		&requestv1.TimestampRequest{
			From: timestamppb.New(from),
			To:   timestamppb.New(to),
		},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrInvalidArgument)
			case codes.NotFound:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrRequestNotFound)
			case codes.Internal:
				return nil, fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return transformRequests(requests.Requests), nil
}

func (ra *requestApi) TimestampAndUserId(
	ctx context.Context,
	from, to time.Time,
	id string,
) ([]*entity.Request, error) {
	const op = "adapter.grpcapi.user.Timestamp"

	requests, err := ra.reqApi.TimestampAndUserId(
		ctx,
		&requestv1.TimestampAndIdRequest{
			From: timestamppb.New(from),
			To:   timestamppb.New(to),
			Id:   id,
		},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrInvalidArgument)
			case codes.NotFound:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrRequestNotFound)
			case codes.Internal:
				return nil, fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return transformRequests(requests.Requests), nil
}

func (ra *requestApi) TimestampAndProxyId(
	ctx context.Context,
	from, to time.Time,
	id string,
) ([]*entity.Request, error) {
	const op = "adapter.grpcapi.user.Timestamp"

	requests, err := ra.reqApi.TimestampAndProxyId(
		ctx,
		&requestv1.TimestampAndIdRequest{
			From: timestamppb.New(from),
			To:   timestamppb.New(to),
			Id:   id,
		},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrInvalidArgument)
			case codes.NotFound:
				return nil, fmt.Errorf("%s: %w", op, adapter.ErrRequestNotFound)
			case codes.Internal:
				return nil, fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return transformRequests(requests.Requests), nil
}

func transformRequests(rs []*requestv1.ProxyRequest) []*entity.Request {
	requests := make([]*entity.Request, len(rs))
	for i := range rs {
		requests[i] = &entity.Request{
			RemoteIP: rs[i].RemoteIp,
			Host:     rs[i].Host,
			Upload:   rs[i].Upload,
			Download: rs[i].Download,
		}
	}

	return requests
}
