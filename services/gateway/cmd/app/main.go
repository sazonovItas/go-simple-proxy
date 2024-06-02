package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/app"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/config"
)

func main() {
	cfg, err := configutils.LoadConfigFromEnv[config.Config]()
	if err != nil {
		panic(err)
	}

	l := logger.NewSlogLogger(
		logger.LogConfig{Environment: cfg.Env, LogLevel: logger.DEBUG},
		os.Stdout,
	).With(slog.String("app", "auth-service"))

	application := app.New(l, cfg)
	defer application.Stop()

	go func() {
		application.HTTPServer.MustRun()
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	application.HTTPServer.Stop()
}
