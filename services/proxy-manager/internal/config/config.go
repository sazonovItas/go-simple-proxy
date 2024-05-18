package config

import (
	"time"
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
	RPCServer  RPCServerConfig  `yaml:"rpc_server"`
	Engine     EngineConfig     `yaml:"engine"`
	Proxies    []ProxyConfig    `yaml:"proxies"`
	ProxyImage ProxyImageConfig `yaml:"proxy_image"`
}

type ProxyConfig struct {
	Env               string        `env:"ENV"`
	ID                string        `env:"PROXY_ID"`
	Port              string        `env:"PORT"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" `
	ReadTimeout       time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT"`
	ShutdownTimeout   time.Duration `env:"SHUTDOWN_TIMEOUT"`
}

type ProxyImageConfig struct {
	Image string `yaml:"image"`
}

type EngineConfig struct {
	CheckTimeout time.Duration `yaml:"check_timeout"`
}

type RPCServerConfig struct{}
