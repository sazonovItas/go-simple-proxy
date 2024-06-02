package engine

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/docker/docker/client"
	"github.com/google/uuid"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
)

const updateTimeout = time.Second * 5

type Engine struct {
	log *slog.Logger

	cli      *client.Client
	proxyImg string

	containers []*ProxyContainer
	stopch     chan struct{}
}

func NewEngine(
	proxyImg string,
	clientConfig DockerClientConfig,
	proxies []config.ProxyConfig,

	l *slog.Logger,
) (*Engine, error) {
	cli, err := client.NewClientWithOpts(
		client.WithHost(clientConfig.Host),
		client.WithTimeout(clientConfig.Timeout),
		client.WithVersion(clientConfig.ApiVersion),
	)
	if err != nil {
		return nil, fmt.Errorf("failed init docker client: %w", err)
	}

	containers := make([]*ProxyContainer, len(proxies))
	for i, cfg := range proxies {
		cfg.ID = uuid.NewString()
		containers[i] = NewProxyContainer(proxyImg, cfg)
	}

	return &Engine{
		log: l,

		cli:        cli,
		proxyImg:   proxyImg,
		containers: containers,
		stopch:     make(chan struct{}),
	}, nil
}

func (e *Engine) Run(ctx context.Context) error {
	const op = "engine.Run"

	l := e.log.With(slog.String("op", op))

	for _, ctr := range e.containers {
		err := e.CreateContainer(ctx, ctr)
		if err != nil {
			l.Error("failed to create container", slogger.Err(err))

			return fmt.Errorf("%s: %w", op, err)
		}

		err = e.StartContainer(ctx, ctr)
		if err != nil {
			l.Error("failed to start container", slogger.Err(err))

			return fmt.Errorf("%s: %w", op, err)
		}
	}

	go e.updateContainerState()

	return nil
}

func (e *Engine) ContainersInfo(ctx context.Context) []*ProxyContainer {
	const op = "engine.ContainersInfo"

	return e.containers
}

func (e *Engine) Shutdown(ctx context.Context) {
	const op = "engine.Shutdown"

	l := e.log.With(slog.String("op", op))

	l.Info("stopping proxy manager engine")

	close(e.stopch)

	for _, ctr := range e.containers {
		err := e.StopContainer(ctx, ctr)
		if err != nil {
			l.Error("failed to stop container", slogger.Err(err))
		}

		err = e.RemoveContainer(ctx, ctr)
		if err != nil {
			l.Error("failed to remove container", slogger.Err(err))
		}
	}
}

func (e *Engine) updateContainerState() {
	for {
		select {
		case <-e.stopch:
		case <-time.After(updateTimeout):
			for _, ctr := range e.containers {
				err := e.ContainerState(context.Background(), ctr)
				if err != nil {
					e.log.Error("failed get info", "container_id", ctr.ID)
				}
			}
		}
	}
}
