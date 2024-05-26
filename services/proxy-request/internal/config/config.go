package config

import "time"

type Config struct {
	Env        string           `yaml:"env"         env:"ENV" env-default:"local"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	RPCServer  GRPCServerConfig `yaml:"rpc_server"`
	Storage    StorageConfig    `yaml:"storage"`
}

type StorageConfig struct {
	Uri       string            `yaml:"uri"          env:"DATABASE_URI"`
	TableName string            `yaml:"table_name"   env:"DB_TABLE_NAME" env-default:"proxy_requests"`
	Conn      StorageConnConfig `yaml:"conn_setting"`
}

type StorageConnConfig struct {
	MaxOpenConns    int           `yaml:"max_open_conns"     env:"DB_MAX_OPEN_CONNS"     env-default:"25"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"  env:"DB_CONN_MAX_LIFETIME"  env-default:"5m"`
	MaxIdleConns    int           `yaml:"max_idle_conns"     env:"DB_IDLE_CONNS"         env-default:"10"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env:"DB_CONN_MAX_IDLE_TIME" env-default:"10m"`
}

type HTTPServerConfig struct {
	Address         string        `yaml:"address"          env:"HTTP_SERVER_ADDRESS"          env-default:":3123"`
	Timeout         time.Duration `yaml:"timeout"          env:"HTTP_SERVER_TIMEOUT"          env-default:"5s"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"     env:"HTTP_SERVER_IDLE_TIMEOUT"     env-default:"30s"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
}

type GRPCServerConfig struct {
	Address string        `yaml:"address" env:"GRPC_SERVER_ADDRESS" env-default:":3223"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_SERVER_TIMEOUT" env-default:"5s"`
}
