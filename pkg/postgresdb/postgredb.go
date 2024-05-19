package postgresdb

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type ConnectionOptions struct {
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

func Connect(ctx context.Context, uri string, opts *ConnectionOptions) (*sqlx.DB, error) {
	const op = "pkg.postgresdb.Connect"

	db, err := sqlx.ConnectContext(ctx, "pgx", uri)
	if err != nil {
		return nil, fmt.Errorf("%s: failed connect to db: %w", op, err)
	}

	db.SetMaxOpenConns(opts.MaxOpenConns)
	db.SetConnMaxLifetime(opts.ConnMaxLifetime)
	db.SetMaxIdleConns(opts.MaxIdleConns)
	db.SetConnMaxIdleTime(opts.ConnMaxIdleTime)

	return db, nil
}
