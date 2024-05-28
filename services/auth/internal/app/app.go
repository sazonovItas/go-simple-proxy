package app

import (
	"context"
	"log/slog"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"

	pguser "github.com/sazonovItas/proxy-manager/services/auth/internal/adapter/pgrepo/user"
	grpcapp "github.com/sazonovItas/proxy-manager/services/auth/internal/app/grpc"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/config"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/lib/hasher"
	authsvc "github.com/sazonovItas/proxy-manager/services/auth/internal/service/auth"
)

type App struct {
	GRPCServer *grpcapp.App

	db *postgresdb.DB
}

func New(
	l *slog.Logger,
	cfg *config.Config,
) *App {
	db, err := postgresdb.Connect(
		context.Background(),
		cfg.Storage.Uri,
		&postgresdb.ConnectionOptions{
			MaxOpenConns:    cfg.Storage.Conn.MaxOpenConns,
			ConnMaxLifetime: cfg.Storage.Conn.ConnMaxLifetime,
			MaxIdleConns:    cfg.Storage.Conn.MaxIdleConns,
			ConnMaxIdleTime: cfg.Storage.Conn.ConnMaxIdleTime,
		},
	)
	if err != nil {
		panic(err)
	}

	hasher := hasher.New(hasher.DefaultCost)
	userRepo := pguser.New(db, cfg.Storage.TableName)

	authSvc := authsvc.New(userRepo, hasher, l, cfg.AuthTokenSecret, cfg.TokenTTL)

	srv := grpcapp.New(&cfg.GRPCServer, l, authSvc)

	return &App{
		GRPCServer: srv,

		db: db,
	}
}

func (a *App) Stop() {
	a.db.Close()
}
