package managercfg

import (
	"time"

	proxycfg "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
)

type Config struct {
	DockerClient DockerClientConfig `yaml:"docker_client"`
	ProxyManager ProxyManagerConfig `yaml:"proxy_manager"`
}

type DockerClientConfig struct {
	Host       string        `yaml:"host"`
	Timeout    time.Duration `yaml:"timeout"`
	ApiVersion string        `yaml:"api_version"`
}

type ProxyManagerConfig struct {
	RPCServer  RPCServerConfig   `yaml:"rpc_server"`
	Engine     EngineConfig      `yaml:"engine"`
	Proxies    []proxycfg.Config `yaml:"proxies"`
	ProxyImage ProxyImageConfig  `yaml:"proxy_image"`
}

type ProxyImageConfig struct {
	Image string `yaml:"image"`
}

type EngineConfig struct {
	CheckTimeout time.Duration `yaml:"check_timeout"`
}

type RPCServerConfig struct{}
