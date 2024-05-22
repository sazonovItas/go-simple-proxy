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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcrequest "github.com/sazonovItas/proxy-manager/services/proxy/internal/adapter/grpc/request"
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

	l := logger.NewSlogLogger(
		logger.LogConfig{Environment: cfg.Env, LogLevel: logger.DEBUG},
		os.Stdout,
	)
	l.Info("config loaded", "config", cfg)

	// init repository
	client, err := grpc.NewClient(":3223", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Error("failed connect to request service", slogger.Err(err))
	}
	defer client.Close()

	requestRepo := grpcrequest.NewRequestRepository(client)

	proxyHandler := proxyhandler.NewProxyHandler(cfg.ID, l, requestRepo)
	handler := middleware.ProxyBasicAuth("proxy")(proxyHandler)
	handler = middleware.Logger(l)(handler)
	handler = middleware.RequestId()(handler)
	handler = middleware.Panic(l)(handler)

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
		l.Info("proxy server started", "address", proxyServer.Addr)
		err := proxyServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			l.Error("server shutdown with error", "error", err.Error())
		}
	}()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer func() {
		cancel()

		if shutdownCtx.Err() != nil && !errors.Is(shutdownCtx.Err(), context.Canceled) {
			l.Warn("proxy shutdown with error", slogger.Err(shutdownCtx.Err()))
		}
	}()

	if err := proxyServer.Shutdown(shutdownCtx); err != nil {
		l.Error("server is shuted down with error", "error", err.Error())
	}

	l.Info("server is shuted down")
}
