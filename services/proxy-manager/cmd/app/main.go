package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/app"
	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
)

const (
	configPathEnv = "CONFIG_PATH"
)

func main() {
	cfg, err := configutils.LoadConfigFromFile[config.Config](os.Getenv(configPathEnv))
	if err != nil {
		log.Fatalf("faild to load proxy manager config: %s", err.Error())
	}

	l := logger.NewSlogLogger(
		logger.LogConfig{Environment: "dev", LogLevel: logger.DEBUG},
		os.Stdout,
	).With(slog.String("app", "proxy-manager-service"))
	l.Info("init config", "config", *cfg)

	application := app.New(l, cfg)
	defer application.Stop()

	go func() {
		application.GRPCServer.MustRun()
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	application.GRPCServer.Stop()
}
