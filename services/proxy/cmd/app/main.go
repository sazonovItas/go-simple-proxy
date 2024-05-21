package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/config"
	proxyhandler "github.com/sazonovItas/proxy-manager/services/proxy/internal/handler"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/handler/middleware"
)

func main() {
	cfg, err := configutils.LoadConfigFromEnv[config.Config]()
	if err != nil {
		log.Fatalf("failed load config with error: %s", err.Error())
		return
	}

	logger := logger.NewSlogLogger(
		logger.LogConfig{Environment: cfg.Env, LogLevel: logger.DEBUG},
		os.Stdout,
	)

	logger.Info("config loaded", "config", cfg)

	proxyHandler := proxyhandler.NewProxyHandler(logger)
	handler := middleware.ProxyBasicAuth("proxy")(proxyHandler)
	handler = middleware.Logger(logger)(handler)
	handler = middleware.RequestId()(handler)
	handler = middleware.Panic(logger)(handler)

	proxyServer := http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("proxy server started", "address", proxyServer.Addr)
		err := proxyServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server shutdown with error", "error", err.Error())
		}
	}()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer func() {
		cancel()

		if shutdownCtx.Err() != nil && !errors.Is(shutdownCtx.Err(), context.Canceled) {
			logger.Warn("proxy shutdown with error", slogger.Err(shutdownCtx.Err()))
		}
	}()

	if err := proxyServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("server is shuted down with error", "error", err.Error())
	}

	logger.Info("server is shuted down")
}
