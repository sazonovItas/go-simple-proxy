package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/config"
	proxyhandler "github.com/sazonovItas/proxy-manager/services/proxy/internal/handler"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/handler/middleware"
)

type App struct {
	log *slog.Logger

	cfg         *config.ProxyConfig
	proxyServer *http.Server
}

func New(
	l *slog.Logger,
	cfg *config.ProxyConfig,
) *App {
	handler := http.Handler(
		proxyhandler.NewProxyHandler(cfg.ID, cfg.Name, cfg.DialTimeout, l, nil, nil),
	)

	// Use middlewares
	handler = middleware.ProxyBasicAuth("")(handler)
	handler = middleware.Logger(l)(handler)
	handler = middleware.Panic(l)(handler)

	proxyServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	return &App{
		log: l,

		cfg:         cfg,
		proxyServer: proxyServer,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.Run"

	a.log.Info("proxy server starting", "address", a.proxyServer.Addr)
	if err := a.proxyServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "app.Stop"

	l := a.log.With(slog.String("op", op))
	l.Info("stopping proxy server")

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
	defer cancel()

	if err := a.proxyServer.Shutdown(ctx); err != nil {
		l.Error("stopping proxy server", slogger.Err(err))
	}
}
