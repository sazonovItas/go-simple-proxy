package httpapp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/gateway/internal/config"
	httphandler "github.com/sazonovItas/proxy-manager/services/gateway/internal/handler/http"
	v1 "github.com/sazonovItas/proxy-manager/services/gateway/internal/handler/http/api/v1"
)

type App struct {
	log *slog.Logger
	cfg *config.HTTPServerConfig

	httpServer *http.Server
	userSvc    v1.UserService
	reqSvc     v1.RequestService
	proxySvc   v1.ProxyService
}

func New(
	l *slog.Logger,
	cfg *config.HTTPServerConfig,
	userSvc v1.UserService,
	reqSvc v1.RequestService,
	proxySvc v1.ProxyService,
) *App {
	router := httphandler.NewRouter(cfg.Timeout, cfg.Port)

	handler := v1.NewHandler(userSvc, reqSvc, proxySvc)
	handler.InitRoutes(router.Group("/api"))

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           router,
		ReadHeaderTimeout: cfg.Timeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	return &App{
		log: l,
		cfg: cfg,

		httpServer: httpServer,
		userSvc:    userSvc,
		reqSvc:     reqSvc,
		proxySvc:   proxySvc,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	a.log.With(slog.String("op", op)).Info("starting http server", "address", a.httpServer.Addr)

	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "httpapp.Stop"

	l := a.log.With(slog.String("op", op))
	l.Info("stopping http server")

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Timeout)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		l.Error("failed shutdown http server", slogger.Err(err))
	}
}
