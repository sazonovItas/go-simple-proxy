package engine

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
)

type Engine struct {
	cli *client.Client

	proxyImg string

	containers []*ProxyContainer
}

func NewEngine(
	proxyImage string,
	clientConfig DockerClientConfig,
) (*Engine, error) {
	cli, err := client.NewClientWithOpts(
		client.WithHost(clientConfig.Host),
		client.WithTimeout(clientConfig.Timeout),
		client.WithVersion(clientConfig.ApiVersion),
	)
	if err != nil {
		return nil, fmt.Errorf("failed init docker client: %w", err)
	}

	return &Engine{
		cli:      cli,
		proxyImg: proxyImage,
	}, nil
}

func (e *Engine) Run(startUpProxies []config.ProxyConfig) (err error) {
	for _, cfg := range startUpProxies {
		ctr := NewProxyContainer("", cfg, e.proxyImg)

		resp, err := e.cli.ContainerCreate(
			context.Background(),
			ctr.Container.ContainerCfg,
			ctr.Container.HostCfg,
			ctr.Container.NetworkCfg,
			ctr.Container.PlatformCfg,
			ctr.Container.Name,
		)
		if err != nil {
			return fmt.Errorf("failed create container: %w", err)
		}

		ctr.ID = resp.ID
		err = e.cli.ContainerStart(context.Background(), ctr.ID, container.StartOptions{})
		if err != nil {
			return fmt.Errorf("failed start container: %w", err)
		}

		stats, err := e.cli.ContainerInspect(context.Background(), ctr.ID)
		if err != nil {
			return fmt.Errorf("failed inspect container: %w", err)
		}

		ctr.ContainerState = stats.ContainerJSONBase.State
		e.containers = append(e.containers, ctr)
	}

	return nil
}

func (e *Engine) Shutdown(ctx context.Context) error {
	for _, ctr := range e.containers {
		err := e.cli.ContainerStop(ctx, ctr.ID, container.StopOptions{Timeout: nil})
		if err != nil {
			return fmt.Errorf("failed stop container: %w", err)
		}

		err = e.cli.ContainerRemove(
			context.Background(),
			ctr.ID,
			container.RemoveOptions{
				Force: true,
			},
		)
		if err != nil {
			return fmt.Errorf("failed remove container: %w", err)
		}
	}

	return nil
}
