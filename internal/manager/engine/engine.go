package engine

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	managercfg "github.com/sazonovItas/go-simple-proxy/internal/config/manager"
	proxycfg "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
	slogger "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl"
)

type Engine struct {
	cli    *client.Client
	logger *slog.Logger

	proxies      []Proxy
	checkTimeout time.Duration

	stopch chan struct{}
}

type Proxy struct {
	proxyID     string
	containerID string
}

func NewEngine(
	dockerCfg managercfg.DockerClientConfig,
	proxyImage managercfg.ProxyImageConfig,
	proxies []proxycfg.Config,
	checkTimeout time.Duration,
	logger *slog.Logger,
) (*Engine, error) {
	cli, err := client.NewClientWithOpts(
		client.WithVersion(dockerCfg.ApiVersion),
		client.WithTimeout(dockerCfg.Timeout),
		client.WithHost(dockerCfg.Host),
	)
	if err != nil {
		return nil, fmt.Errorf("failed init docker client: %w", err)
	}

	containers := make([]Proxy, 0, len(proxies))
	for range proxies {
		proxyID := uuid.New().String()

		containerID, err := cli.ContainerCreate(
			context.Background(),
			&container.Config{
				Image:        proxyImage.Image,
				ExposedPorts: nat.PortSet{"8123": struct{}{}},
				Env:          []string{},
			},
			&container.HostConfig{},
			&network.NetworkingConfig{},
			&v1.Platform{},
			proxyID,
		)
		if err != nil {
			logger.Error("create container", slogger.Err(err))
		}

		containers = append(containers, Proxy{
			containerID: containerID.ID,
			proxyID:     proxyID,
		})
	}

	return &Engine{
		cli:    cli,
		logger: logger,

		checkTimeout: checkTimeout,
		proxies:      containers,

		stopch: make(chan struct{}),
	}, nil
}

func (e *Engine) Run() {
	for {
		select {
		case <-e.stopch:
			return
		case <-time.After(e.checkTimeout):
		}
	}
}

func (e *Engine) Shutdown() {
	e.stopch <- struct{}{}
}
