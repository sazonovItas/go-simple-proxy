package config

import "time"

type Config struct {
	Env      string         `yaml:"env"      env:"ENV" env-default:"local"`
	Proxy    ProxyConfig    `yaml:"proxy"`
	Services ServicesConfig `yaml:"services"`
}

type ProxyConfig struct {
	ID                string        `yaml:"proxy_id"            env:"PROXY_ID"            env-default:""`
	Name              string        `yaml:"name"                env:"PROXY_NAME"          env-default:""`
	Port              int           `yaml:"port"                env:"PORT"                env-default:"8123"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"READ_HEADER_TIMEOUT" env-default:"5s"`
	DialTimeout       time.Duration `yaml:"dial_timeout"        env:"DIAL_TIMEOUT"        env-default:"5s"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout"    env:"SHUTDOWN_TIMEOUT"    env-default:"10s"`
}

type ServicesConfig struct {
	RequestServiceAddr string `yaml:"request_service_address" env:"REQUEST_SERVICE_ADDR" env-default:":3223"`
	UserServiceAddr    string `yaml:"user_service_addr"       env:"USER_SERVICE_ADDR"    env-default:":3224"`
}
