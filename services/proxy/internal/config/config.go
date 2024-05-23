package config

import "time"

type Config struct {
	Env               string        `yaml:"env"                 env:"ENV"                 env-default:"local"`
	ID                string        `yaml:"proxy_id"            env:"PROXY_ID"            env-default:""`
	Port              string        `yaml:"port"                env:"PORT"                env-default:"8123"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"READ_HEADER_TIMEOUT" env-default:"1s"`
	ReadTimeout       time.Duration `yaml:"read_timeout"        env:"READ_TIMEOUT"        env-default:"30s"`
	WriteTimeout      time.Duration `yaml:"write_timeout"       env:"WRITE_TIMEOUT"       env-default:"30s"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout"    env:"SHUTDOWN_TIMEOUT"    env-default:"10s"`

	RequestServiceAddr string `yaml:"request_service_address" env:"REQUEST_SERVICE_ADDR"`
	UserServiceAddr    string `yaml:"user_service_addr"       env:"USER_SERVICE_ADDR"`
}
