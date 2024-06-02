package manager

import (
	"context"

	"google.golang.org/grpc"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/entity"
	managerv1 "github.com/sazonovItas/proxy-manager/services/proxy-manager/pkg/pb/manager/v1"
)

type ManagerService interface {
	Info(ctx context.Context) []*entity.ProxyState
}

type managerHandler struct {
	managerSvc ManagerService

	managerv1.UnimplementedProxyManagerServer
}

func Register(srv *grpc.Server, handler *managerHandler) {
	managerv1.RegisterProxyManagerServer(srv, handler)
}

func New(managerSvc ManagerService) *managerHandler {
	return &managerHandler{
		managerSvc: managerSvc,
	}
}

func (mh *managerHandler) Info(
	ctx context.Context,
	r *managerv1.InfoRequest,
) (*managerv1.InfoResponse, error) {
	states := mh.managerSvc.Info(ctx)

	info := make([]*managerv1.ProxyState, 0, len(states))
	for _, state := range states {
		info = append(info, &managerv1.ProxyState{
			Id:        state.ID,
			Status:    state.Status,
			Address:   state.Address,
			StartedAt: state.StartedAt,
		})
	}

	return &managerv1.InfoResponse{Info: info}, nil
}
