package engine

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
	reflectutils "github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/utils"
)

const (
	ProxyRestartPolicy = container.RestartPolicyOnFailure
	// ProxyRestartCount is amount of time for restarting proxy
	ProxyRestartCount = 3
	// NetworkMode is network mode for proxy
	ProxyNetworkMode = "host"
)

type ProxyContainer struct {
	ID             string
	ProxyID        string
	ContainerState *types.ContainerState
	Container      *ContainerConfig
}

func NewProxyContainer(
	containerName string,
	proxyCfg *config.ProxyConfig,
	proxyImage string,
) *ProxyContainer {
	return &ProxyContainer{
		ProxyID:        uuid.NewString(),
		ContainerState: &types.ContainerState{},
		Container: &ContainerConfig{
			Name: containerName,
			ContainerCfg: &container.Config{
				Env:          reflectutils.StructToEnv(proxyCfg),
				ExposedPorts: nat.PortSet{nat.Port(proxyCfg.Port): struct{}{}},
				Image:        proxyImage,
			},
			HostCfg: &container.HostConfig{
				NetworkMode: ProxyNetworkMode,
				RestartPolicy: container.RestartPolicy{
					Name:              ProxyRestartPolicy,
					MaximumRetryCount: ProxyRestartCount,
				},
			},
			NetworkCfg:  &network.NetworkingConfig{},
			PlatformCfg: &v1.Platform{},
		},
	}
}
