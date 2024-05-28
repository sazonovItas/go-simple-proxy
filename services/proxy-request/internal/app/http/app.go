package httpapp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy-request/internal/config"
	router "github.com/sazonovItas/proxy-manager/services/proxy-request/internal/handler/http"
	httpv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/internal/handler/http/api/v1"
)

type App struct {
	log *slog.Logger

	cfg        *config.HTTPServerConfig
	httpServer *http.Server
}

func New(cfg *config.HTTPServerConfig, l *slog.Logger, requestUsc httpv1.RequestUsecase) *App {
	router := router.New(cfg.Timeout, true)

	handler := httpv1.NewHandler(requestUsc)
	handler.Init(router.Group("/api"))

	srv := &http.Server{
		Handler:     router,
		Addr:        cfg.Address,
		IdleTimeout: cfg.IdleTimeout,
	}

	return &App{
		log: l,

		cfg:        cfg,
		httpServer: srv,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	a.log.With(slog.String("op", op)).
		Info("http server started", slog.String("address", a.httpServer.Addr))

	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "httpapp.Stop"

	l := a.log.With(slog.String("op", op))
	l.Info("stopping http server")

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		l.Error("something goes wrong while http server stopping", slogger.Err(err))
	}
}
