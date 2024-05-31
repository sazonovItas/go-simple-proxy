package config

import "time"

type Config struct {
	Env        string           `yaml:"env"        env:"ENV" env-default:"local"`
	GRPCServer GRPCServerConfig `yaml:"rpc_server"`
	Storage    StorageConfig    `yaml:"storage"`
}

type StorageConfig struct {
	Uri       string            `yaml:"uri"          env:"DATABASE_URI"`
	TableName string            `yaml:"table_name"   env:"DB_TABLE_NAME" env-default:"proxy_requests"`
	Conn      StorageConnConfig `yaml:"conn_setting"`
}

type StorageConnConfig struct {
	MaxOpenConns    int           `yaml:"max_open_conns"     env:"DB_MAX_OPEN_CONNS"     env-default:"20"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"  env:"DB_CONN_MAX_LIFETIME"  env-default:"5m"`
	MaxIdleConns    int           `yaml:"max_idle_conns"     env:"DB_IDLE_CONNS"         env-default:"5"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env:"DB_CONN_MAX_IDLE_TIME" env-default:"10m"`
}

type GRPCServerConfig struct {
	Port    int           `yaml:"port"    env:"GRPC_SERVER_PORT"    env-default:"3223"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_SERVER_TIMEOUT" env-default:"5s"`
}
