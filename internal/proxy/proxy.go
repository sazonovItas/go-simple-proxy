package proxy

import (
	"context"
	"errors"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	proxycfg "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
	proxy "github.com/sazonovItas/go-simple-proxy/internal/proxy/handler"
	"github.com/sazonovItas/go-simple-proxy/internal/proxy/handler/middleware"
	slogger "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl"
)

const (
	local       = "local"
	development = "dev"
	production  = "prod"
)

func Run() {
	cfg, err := configutils.LoadCfgFromEnv[proxycfg.Config]()
	if err != nil {
		log.Fatalf("failed load config with error: %s", err.Error())
		return
	}

	logger := InitLogger(cfg.Env, os.Stdout)

	logger.Info("config loaded", "config", cfg)

	proxyHandler := proxy.NewProxyHandler(logger)
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
		logger.Info("proxy server started", "address", cfg.Port)
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

func InitLogger(env string, out io.Writer) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case development:
		logger = slogger.NewPrettyLogger(slog.LevelInfo, out)
	case production:
		logger = slogger.NewPrettyLogger(slog.LevelWarn, out)
	default:
		logger = slogger.NewPrettyLogger(slog.LevelDebug, out)
	}

	return logger
}
