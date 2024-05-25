package grpcapp

import (
	"log/slog"

	"github.com/sazonovItas/proxy-manager/proxy-request/internal/config"
)

type App struct {
	l *slog.Logger

	cfg *config.GRPCServerConfig
}

func New() *App {
	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}
