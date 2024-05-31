package app

import (
	"context"
	"log/slog"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"

	pgrequest "github.com/sazonovItas/proxy-manager/services/proxy-request/internal/adapter/pgrepo/request"
	grpcapp "github.com/sazonovItas/proxy-manager/services/proxy-request/internal/app/grpc"
	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/config"
	requestusc "github.com/sazonovItas/proxy-manager/services/proxy-request/internal/usecase/request"
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

	requestRepo := pgrequest.New(cfg.Storage.TableName, db)
	requestUsc := requestusc.New(requestRepo)

	gRPCServer := grpcapp.New(&cfg.GRPCServer, l, requestUsc)

	return &App{
		GRPCServer: gRPCServer,

		db: db,
	}
}

func (a *App) Stop() {
	a.db.Close()
}
