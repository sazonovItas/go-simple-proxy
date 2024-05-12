package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	configproxy "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
	proxy "github.com/sazonovItas/go-simple-proxy/internal/proxy/handler"
	"github.com/sazonovItas/go-simple-proxy/internal/proxy/handler/middleware"
	slogger "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl"
)

func main() {
	logger := slogger.NewPrettyLogger(slog.LevelDebug, os.Stdout)

	cfg, err := configutils.LoadCfgFromFile[configproxy.Config](
		"./config/proxy/" + configutils.GetEnv() + ".yaml",
	)
	if err != nil {
		logger.Error("failed load config from file", "error", err.Error())
		return
	}
	logger.Info("config loaded", "config", cfg)

	proxyHandler := proxy.NewProxyHandler(logger)
	handler := middleware.ProxyBasicAuth("proxy")(proxyHandler)
	handler = middleware.Logger(logger)(handler)
	handler = middleware.RequestId()(handler)

	proxyServer := http.Server{
		Addr:              cfg.Proxy.Host + ":" + strconv.Itoa(cfg.Proxy.Port),
		Handler:           handler,
		ReadHeaderTimeout: cfg.Proxy.ReadHeaderTimeout,
		IdleTimeout:       cfg.Proxy.IdleTImeout,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("proxy server started", "host", cfg.Proxy.Host, "port", cfg.Proxy.Port)
		err := proxyServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server shutdown with error", "error", err.Error())
		}
	}()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Proxy.ShutdownTimeout)
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
