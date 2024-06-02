package managersvc

import (
	"context"
	"fmt"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/engine"
	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/entity"
)

type containerRepository interface {
	ContainersInfo(ctx context.Context) []*engine.ProxyContainer
}

type managerService struct {
	host    string
	ctrRepo containerRepository
}

func New(host string, ctrRepo containerRepository) *managerService {
	return &managerService{
		host:    host,
		ctrRepo: ctrRepo,
	}
}

func (cs *managerService) Info(ctx context.Context) []*entity.ProxyState {
	containers := cs.ctrRepo.ContainersInfo(ctx)

	states := make([]*entity.ProxyState, len(containers))
	for i := range containers {
		states[i] = &entity.ProxyState{
			ID:        containers[i].ProxyID,
			Status:    containers[i].ContainerState.Status,
			Address:   fmt.Sprintf("http://%s:%d", cs.host, containers[i].ProxyPort),
			StartedAt: containers[i].ContainerState.StartedAt,
		}
	}

	return states
}
