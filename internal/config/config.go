package config

import configproxy "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"

type Config struct {
	Proxy configproxy.Proxy `yaml:"proxy"`
}
