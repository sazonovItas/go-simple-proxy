package app

import (
	"context"
	"log/slog"

	grpcapp "github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/app/grpc"
	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/engine"
	managersvc "github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/service/manager"
)

type App struct {
	GRPCServer *grpcapp.App

	e   *engine.Engine
	cfg *config.Config
}

func New(l *slog.Logger, cfg *config.Config) *App {
	e, err := engine.NewEngine(
		cfg.ProxyManager.ProxyImage.Image,
		engine.DockerClientConfig{
			ApiVersion: cfg.DockerClient.ApiVersion,
			Host:       cfg.DockerClient.Host,
			Timeout:    cfg.DockerClient.Timeout,
		},
		cfg.ProxyManager.Proxies,
		l,
	)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ProxyManager.StartTimeout)
	defer cancel()

	err = e.Run(ctx)
	if err != nil {
		e.Shutdown(context.Background())
		panic(err)
	}

	gRPCServer := grpcapp.New(&cfg.GRPCServer, l, managersvc.New(cfg.ProxyManager.Host, e))

	return &App{
		e:   e,
		cfg: cfg,

		GRPCServer: gRPCServer,
	}
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ProxyManager.ShutdownTimeout)
	defer cancel()

	a.e.Shutdown(ctx)
}
