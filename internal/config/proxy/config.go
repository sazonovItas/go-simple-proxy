package configproxy

import "time"

type Config struct {
	Env               string        `yaml:"env"                 env:"ENV"`
	Host              string        `yaml:"host"                env:"HOST"`
	Port              int           `yaml:"port"                env:"PORT"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"READ_HEADER_TIMEOUT"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"        env:"IDLE_TIMEOUT"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout"    env:"SHUTDOWN_TIMEOUT"`
}
