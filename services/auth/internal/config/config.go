package config

import "time"

// TODO: Delete default auth token secret
type Config struct {
	Env             string           `yaml:"env"               env:"ENV"               env-default:"local"`
	TokenTTL        time.Duration    `yaml:"token_ttl"         env:"TOKEN_TTL"         env-default:"30m"`
	AuthTokenSecret string           `yaml:"auth_token_secret" env:"AUTH_TOKEN_SECRET" env-default:"AUTH_SERCRET"`
	GRPCServer      GRPCServerConfig `yaml:"grpc_server"`
	Storage         StorageConfig    `yaml:"storage"`
}

type StorageConfig struct {
	Uri       string            `yaml:"uri"          env:"DATABASE_URI"`
	TableName string            `yaml:"table_name"   env:"DB_TABLE_NAME" env-default:"proxy_users"`
	Conn      StorageConnConfig `yaml:"conn_setting"`
}

type StorageConnConfig struct {
	MaxOpenConns    int           `yaml:"max_open_conns"     env:"DB_MAX_OPEN_CONNS"     env-default:"20"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"  env:"DB_CONN_MAX_LIFETIME"  env-default:"5m"`
	MaxIdleConns    int           `yaml:"max_idle_conns"     env:"DB_IDLE_CONNS"         env-default:"5"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env:"DB_CONN_MAX_IDLE_TIME" env-default:"10m"`
}

type GRPCServerConfig struct {
	Address string        `yaml:"address" env:"GRPC_SERVER_ADDRESS" env-default:":3224"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_SERVER_TIMEOUT" env-default:"5s"`
}
