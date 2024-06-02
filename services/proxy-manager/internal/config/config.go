package config

import (
	"time"
)

type Config struct {
	GRPCServer   GRPCServerConfig   `yaml:"grpc_server"`
	DockerClient DockerClientConfig `yaml:"docker_client"`
	ProxyManager ProxyManagerConfig `yaml:"proxy_manager"`
}

type DockerClientConfig struct {
	Host       string        `yaml:"host"`
	Timeout    time.Duration `yaml:"timeout"`
	ApiVersion string        `yaml:"api_version"`
}

type ProxyManagerConfig struct {
	Host            string           `yaml:"host"`
	Proxies         []ProxyConfig    `yaml:"proxy"`
	ProxyImage      ProxyImageConfig `yaml:"proxy_image"`
	ShutdownTimeout time.Duration    `yaml:"shutdown_timeout"`
	StartTimeout    time.Duration    `yaml:"start_timeout"`
}

type ProxyConfig struct {
	Env                string        `env:"ENV"                  json:"env,omitempty"                     yaml:"env"`
	ID                 string        `env:"PROXY_ID"             json:"id,omitempty"                      yaml:"id"`
	Port               int           `env:"PORT"                 json:"port,omitempty"                    yaml:"port"`
	ReadHeaderTimeout  time.Duration `env:"READ_HEADER_TIMEOUT"  json:"read_header_timeout,omitempty"     yaml:"read_header_timeout"`
	DialTimeout        time.Duration `env:"DIAL_TIMEOUT"         json:"dial_timeout,omitempty"            yaml:"dial_timeout"`
	ShutdownTimeout    time.Duration `env:"SHUTDOWN_TIMEOUT"     json:"shutdown_timeout,omitempty"        yaml:"shutdown_timeout"`
	RequestServiceAddr string        `env:"REQUEST_SERVICE_ADDR" json:"request_service_address,omitempty" yaml:"request_service_addr"`
	UserServiceAddr    string        `env:"USER_SERVICE_ADDR"    json:"user_service_address,omitempty"    yaml:"user_service_addr"`
}

type ProxyImageConfig struct {
	Image string `yaml:"image"`
}

type GRPCServerConfig struct {
	Port    int           `yaml:"port"    env:"GRPC_SERVER_PORT"     env-default:"3225"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_SERVIER_TIMEOUT" env-default:"10s"`
}
