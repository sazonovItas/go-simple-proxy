package main

import (
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	configproxy "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
	proxy "github.com/sazonovItas/go-simple-proxy/internal/proxy/handler"
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

	proxyHandler := proxy.NewProxyHandler(logger, cfg.Proxy.BlockList)

	proxyServer := http.Server{
		Addr:              cfg.Proxy.Address,
		ReadHeaderTimeout: cfg.Proxy.ReadHeaderTimeout,
		Handler:           http.HandlerFunc(proxyHandler.ServeHTTP),
		TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		logger.Info("proxy server started", "address", cfg.Proxy.Address)
		err := proxyServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server shutdown with error", "error", err.Error())
		}
	}()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Proxy.ShutdownTimeout)
	defer func() {
		cancel()

		if !errors.Is(shutdownCtx.Err(), context.Canceled) {
			logger.Warn("proxy shutdown with error", slogger.Err(err))
		}
	}()

	if err := proxyServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("server is shuted down with error", "error", err.Error())
	}

	logger.Info("server is shuted down")
}
