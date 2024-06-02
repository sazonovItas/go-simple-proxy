package engine

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/google/uuid"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
	reflectutils "github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/utils"
)

type ProxyContainer struct {
	ID             string
	ProxyID        string
	ProxyPort      int
	ContainerState *types.ContainerState
	Container      *ContainerConfig
}

func NewProxyContainer(
	proxyImg string,
	proxyCfg config.ProxyConfig,
) *ProxyContainer {
	return &ProxyContainer{
		ProxyID:        uuid.NewString(),
		ProxyPort:      proxyCfg.Port,
		ContainerState: &types.ContainerState{},
		Container: &ContainerConfig{
			Container: &container.Config{
				Env:   reflectutils.StructToEnv(proxyCfg),
				Image: proxyImg,
			},
			Host: &container.HostConfig{
				NetworkMode: container.NetworkMode("host"),
				RestartPolicy: container.RestartPolicy{
					Name: container.RestartPolicyAlways,
				},
			},
			Network:  &network.NetworkingConfig{},
			Platform: &v1.Platform{},
		},
	}
}

func (e *Engine) ContainerState(
	ctx context.Context,
	ctr *ProxyContainer,
) error {
	const op = "engine.ContainerState"

	stats, err := e.cli.ContainerInspect(ctx, ctr.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	ctr.ContainerState = stats.State

	return nil
}

func (e *Engine) CreateContainer(ctx context.Context, ctr *ProxyContainer) error {
	const op = "engine.CreateContainer"

	resp, err := e.cli.ContainerCreate(
		ctx,
		ctr.Container.Container,
		ctr.Container.Host,
		ctr.Container.Network,
		ctr.Container.Platform,
		namesgenerator.GetRandomName(10)+fmt.Sprintf("-%d", ctr.ProxyPort),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	ctr.ID = resp.ID

	stats, err := e.cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	ctr.ContainerState = stats.ContainerJSONBase.State

	return nil
}

func (e *Engine) StartContainer(ctx context.Context, ctr *ProxyContainer) error {
	const op = "engine.StartContainer"

	err := e.cli.ContainerStart(ctx, ctr.ID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (e *Engine) StopContainer(ctx context.Context, ctr *ProxyContainer) error {
	const op = "engine.StopContainer"

	err := e.cli.ContainerStop(ctx, ctr.ID, container.StopOptions{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (e *Engine) RemoveContainer(ctx context.Context, ctr *ProxyContainer) error {
	const op = "engine.RemoveContainer"

	err := e.cli.ContainerRemove(ctx, ctr.ID, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
