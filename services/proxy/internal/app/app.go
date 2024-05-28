package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcrequest "github.com/sazonovItas/proxy-manager/services/proxy/internal/adapter/grpc/request"
	grpcuser "github.com/sazonovItas/proxy-manager/services/proxy/internal/adapter/grpc/user"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/config"
	proxyhandler "github.com/sazonovItas/proxy-manager/services/proxy/internal/handler"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/handler/middleware"
	proxysvc "github.com/sazonovItas/proxy-manager/services/proxy/internal/service/proxy"
)

type App struct {
	log *slog.Logger

	cfg         *config.Config
	proxyServer *http.Server

	cliUser    *grpc.ClientConn
	cliRequest *grpc.ClientConn
}

func New(
	l *slog.Logger,
	cfg *config.Config,
) *App {
	cliRequest, err := grpc.NewClient(
		cfg.Services.RequestServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	cliUser, err := grpc.NewClient(
		cfg.Services.UserServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	requestRepo := grpcrequest.New(cliRequest)
	userRepo := grpcuser.New(nil)

	if cfg.Proxy.ID == "" {
		cfg.Proxy.ID = uuid.NewString()
	}

	handler := http.Handler(
		proxyhandler.New(
			cfg.Proxy.ID,
			cfg.Proxy.Name,
			cfg.Proxy.DialTimeout,
			l,
			proxysvc.New(requestRepo, userRepo),
		),
	)

	// Use middlewares
	handler = middleware.ProxyBasicAuth("")(handler)
	handler = middleware.Logger(l)(handler)
	handler = middleware.Panic(l)(handler)

	proxyServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Proxy.Port),
		Handler:           handler,
		ReadHeaderTimeout: cfg.Proxy.ReadHeaderTimeout,
	}

	return &App{
		log: l,

		cfg:         cfg,
		proxyServer: proxyServer,

		cliUser:    cliUser,
		cliRequest: cliRequest,
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

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Proxy.ShutdownTimeout)
	defer cancel()

	if err := a.proxyServer.Shutdown(ctx); err != nil {
		l.Error("stopping proxy server", slogger.Err(err))
	}

	a.cliRequest.Close()
	a.cliUser.Close()
}
