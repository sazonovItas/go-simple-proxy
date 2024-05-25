package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	requestrepo "github.com/sazonovItas/proxy-manager/services/proxy/internal/adapter/grpc/request"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/config"
	proxyhandler "github.com/sazonovItas/proxy-manager/services/proxy/internal/handler"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/handler/middleware"
	proxysvc "github.com/sazonovItas/proxy-manager/services/proxy/internal/service/proxy"
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
	client, err := grpc.NewClient(
		cfg.Services.RequestServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Error("failed connect to request service", slogger.Err(err))
	}
	defer client.Close()

	requestRepo := requestrepo.NewRequestRepository(client)
	proxysvc.New(requestRepo, nil)

	proxyHandler := proxyhandler.NewProxyHandler(
		uuid.NewString(),
		uuid.NewString(),
		cfg.Proxy.DialTimeout,
		l,
		requestRepo,
		nil,
	)
	handler := middleware.ProxyBasicAuth("proxy")(proxyHandler)
	handler = middleware.Logger(l)(handler)
	handler = middleware.RequestId()(handler)
	handler = middleware.Panic(l)(handler)

	proxyServer := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Proxy.Port),
		Handler:           handler,
		ReadHeaderTimeout: cfg.Proxy.ReadHeaderTimeout,
	}

	go func() {
		l.Info("proxy server started", "address", proxyServer.Addr)
		err := proxyServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			l.Error("server shutdown with error", "error", err.Error())
		}
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Proxy.ShutdownTimeout)
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
