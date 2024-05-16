package proxycfg

import (
	"time"
)

type Config struct {
	Env               string        `yaml:"env"                 env:"ENV"                 env-default:"local"`
	ID                string        `yaml:"proxy_id"            env:"PROXY_ID"            env-default:"00000000-0000-0000-0000-00000000"`
	Host              string        `yaml:"host"                env:"HOST"                env-default:"0.0.0.0"`
	Port              string        `yaml:"port"                env:"PORT"                env-default:"8123"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"READ_HEADER_TIMEOUT" env-default:"1s"`
	ReadTimeout       time.Duration `yaml:"read_timeout"        env:"READ_TIMEOUT"        env-default:"30s"`
	WriteTimeout      time.Duration `yaml:"write_timeout"       env:"WRITE_TIMEOUT"       env-default:"30s"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout"    env:"SHUTDOWN_TIMEOUT"    env-default:"10s"`
}
