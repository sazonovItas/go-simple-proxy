package config

import "time"

type Config struct {
	Env        string           `yaml:"env"         env:"ENV" env-default:"local"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	RPCServer  RPCServerConfig  `yaml:"rpc_server"`
	Storage    StorageConfig    `yaml:"storage"`
}

type StorageConfig struct {
	Uri string `yaml:"uri" env:"DATABASE_URI"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address"      env:"HTTP_SERVER_ADDRESS"      env-default:":3123"`
	Timeout     time.Duration `yaml:"timeout"      env:"HTTP_SERVER_TIMEOUT"      env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"30s"`
}

type RPCServerConfig struct {
	Address      string        `yaml:"address"       env:"RPC_SERVER_ADDRESS"       env-default:":3223"`
	ReadTimeout  time.Duration `yaml:"read_timeout"  env:"RPC_SERVER_READ_TIMEOUT"  env-default:"5s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"RPC_SERVER_WRITE_TIMEOUT" env-default:"5s"`
}
