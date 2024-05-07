package configproxy

import "time"

type Config struct {
	Proxy Proxy `yaml:"proxy"`
}

type Proxy struct {
	Address           string        `yaml:"address"`
	ShutdownTimeout   time.Duration `yaml:"shutdown_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	Secrets           ProxySecrets  `yaml:"secrets"`
	BlockList         []string      `yaml:"block-list"`
}

type ProxySecrets struct {
	Key  string `yaml:"key"`
	Cert string `yaml:"cert"`
}
