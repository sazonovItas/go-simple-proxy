package app

import (
	"context"
	"log/slog"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"

	pgrequest "github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter/pgrepo/request"
	grpcapp "github.com/sazonovItas/proxy-manager/proxy-request/internal/app/grpc"
	httpapp "github.com/sazonovItas/proxy-manager/proxy-request/internal/app/http"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/config"
)

type App struct {
	GRPCServer *grpcapp.App
	HTTPServer *httpapp.App

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

	gRPCServer := grpcapp.New(&cfg.GRPCServer, l, requestRepo)
	httpServer := httpapp.New(&cfg.HTTPServer, l, requestRepo)

	return &App{
		GRPCServer: gRPCServer,
		HTTPServer: httpServer,

		db: db,
	}
}

func (a *App) Stop() {
	a.db.Close()
}
