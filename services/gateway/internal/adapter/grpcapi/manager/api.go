package grpcmanagerapi

import (
	"context"
	"fmt"

	managerv1 "github.com/sazonovItas/proxy-manager/services/proxy-manager/pkg/pb/manager/v1"
	"google.golang.org/grpc"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/entity"
)

type grpcManagerApi interface {
	Info(
		ctx context.Context,
		in *managerv1.InfoRequest,
		opts ...grpc.CallOption,
	) (*managerv1.InfoResponse, error)
}

type managerApi struct {
	mngApi grpcManagerApi
}

func New(mngApi grpcManagerApi) *managerApi {
	return &managerApi{
		mngApi: mngApi,
	}
}

func (ma *managerApi) ProxyInfo(ctx context.Context) ([]*entity.Proxy, error) {
	const op = "adapter.grpcapi.manager.ProxyInfo"

	resp, err := ma.mngApi.Info(ctx, &managerv1.InfoRequest{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return transformProxyStates(resp.Info), nil
}

func transformProxyStates(proxyStates []*managerv1.ProxyState) []*entity.Proxy {
	proxies := make([]*entity.Proxy, len(proxyStates))
	for i := range proxyStates {
		proxies[i] = &entity.Proxy{
			ID:        proxyStates[i].GetId(),
			Status:    proxyStates[i].GetStatus(),
			Address:   proxyStates[i].GetAddress(),
			StartedAt: proxyStates[i].GetStartedAt(),
		}
	}

	return proxies
}
