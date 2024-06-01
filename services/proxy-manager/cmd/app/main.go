package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/config"
	"github.com/sazonovItas/proxy-manager/services/proxy-manager/internal/engine"
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
	)
	l.Info("init config", "config", *cfg)

	engine, err := engine.NewEngine(cfg.ProxyManager.ProxyImage.Image, engine.DockerClientConfig{
		ApiVersion: cfg.DockerClient.ApiVersion,
		Timeout:    cfg.DockerClient.Timeout,
		Host:       cfg.DockerClient.Host,
	}, cfg.ProxyManager.Proxies, l)
	if err != nil {
		l.Error("failed init engine", slogger.Err(err))
		return
	}

	err = engine.Run(context.Background())
	if err != nil {
		l.Error("failed to run engine", slogger.Err(err))
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		cfg.ProxyManager.ShutdownTimeout,
	)
	defer cancel()

	engine.Shutdown(shutdownCtx)

	l.Info("proxy manager stopped")
}
