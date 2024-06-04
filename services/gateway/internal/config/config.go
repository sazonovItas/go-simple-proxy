package config

import "time"

type Config struct {
	Env        string           `yaml:"env"         env:"ENV" env-default:"dev"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	Services   ServicesConfig   `yaml:"services"`
}

type HTTPServerConfig struct {
	Port            int           `yaml:"port"             env:"HTTP_SERVER_PORT"             env-default:"3030"`
	Timeout         time.Duration `yaml:"timeout"          env:"HTTP_SERVER_TIMEOUT"          env-default:"10s"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"     env:"HTTP_SERVER_IDLE_TIMEOUT"     env-default:"60s"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" env-default:"10s"`
	Origin          string        `yaml:"origin"           env:"HTTP_SERVER_ORIGIN"`
}

type ServicesConfig struct {
	RequestSvcAddr string `yaml:"request_svc_addr" env:"REQUEST_SERVICE_ADDR" env-default:":3223"`
	AuthSvcAddr    string `yaml:"auth_svc_addr"    env:"AUTH_SERVICE_ADDR"    env-default:":3224"`
	ManagerSvcAddr string `yaml:"manager_svc_addr" env:"MANAGER_SERVICE_ADDR" env-default:":3225"`
}
