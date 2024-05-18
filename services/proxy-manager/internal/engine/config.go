package engine

import (
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DockerClientConfig struct {
	Host       string
	Timeout    time.Duration
	ApiVersion string
}

type ContainerConfig struct {
	Name         string
	ContainerCfg *container.Config
	HostCfg      *container.HostConfig
	NetworkCfg   *network.NetworkingConfig
	PlatformCfg  *v1.Platform
}
