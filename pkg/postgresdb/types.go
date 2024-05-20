package postgresdb

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ConnectionOptions struct {
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type DB struct {
	*sqlx.DB
}
