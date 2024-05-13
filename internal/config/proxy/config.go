package proxycfg

import "time"

type Config struct {
	Env               string        `yaml:"env"                 env:"ENV"`
	ID                string        `yaml:"proxy_id"            env:"PROXY_ID"`
	Address           string        `yaml:"address"             env:"ADDRESS"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" env:"READ_HEADER_TIMEOUT"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"        env:"IDLE_TIMEOUT"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout"    env:"SHUTDOWN_TIMEOUT"`
}
