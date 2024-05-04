package app

import (
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sazonovItas/go-simple-proxy/internal/config"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
	proxy "github.com/sazonovItas/go-simple-proxy/internal/proxy/handler"
	slogger "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl"
)

func Run() {
	logger := slogger.NewPrettyLogger(slog.LevelDebug, os.Stdout)

	cfg, err := configutils.LoadCfgFromFile[config.Config](
		"./config/" + configutils.GetEnv() + ".yaml",
	)
	if err != nil {
		logger.Error("failed load config from file", "error", err.Error())
		return
	}
	logger.Info("config loaded", "config", cfg)

	proxyHandler := proxy.NewProxyHandler(logger, cfg.Proxy.BlockList)

	proxyServer := http.Server{
		Addr:         cfg.Proxy.Address,
		Handler:      http.HandlerFunc(proxyHandler.ServeHTTP),
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
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

	if err := proxyServer.Shutdown(context.Background()); err != nil {
		logger.Error("server is shuted down with error", "error", err.Error())
	}

	logger.Info("server is shuted down")
}
