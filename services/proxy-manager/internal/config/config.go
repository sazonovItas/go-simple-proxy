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
	Env               string        `env:"ENV"                 json:"env"`
	ID                string        `env:"PROXY_ID"            json:"id"`
	Port              string        `env:"PORT"                json:"port"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" json:"read_header_timeout"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT"        json:"read_timeout"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT"       json:"write_timeout"`
	ShutdownTimeout   time.Duration `env:"SHUTDOWN_TIMEOUT"    json:"shutdown_timeout"`

	ProxyUserCacheTimeout   time.Duration `env:"PROXY_USER_CACHE_TIMEOUT"      json:"proxy_user_cache_timeout"`
	ProxyUserServiceAddr    string        `env:"PROXY_USER_SERVICE_ADDRESS"    json:"proxy_user_service_addr"`
	ProxyRequestServiceAddr string        `env:"PROXY_REQUEST_SERVICE_ADDRESS" json:"proxy_request_service_addr"`
}

type ProxyContainerConfig struct {
	Image           string `yaml:"image"            env:"PROXY_IMAGE"`
	ServicesNetwork string `yaml:"services_network" env:"SERVICE_NETWORK"`
}

type ProxyImageConfig struct {
	Image string `yaml:"image"`
}

type EngineConfig struct {
	CheckTimeout time.Duration `yaml:"check_timeout"`
}

type RPCServerConfig struct{}
