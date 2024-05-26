package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"
	requestv1 "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb/request/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcrequest "github.com/sazonovItas/proxy-manager/services/proxy/internal/adapter/grpc/request"
	grpcuser "github.com/sazonovItas/proxy-manager/services/proxy/internal/adapter/grpc/user"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/app"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/config"
	proxysvc "github.com/sazonovItas/proxy-manager/services/proxy/internal/service/proxy"
)

func main() {
	cfg, err := configutils.LoadConfigFromEnv[config.Config]()
	if err != nil {
		panic(err)
	}

	l := logger.NewSlogLogger(
		logger.LogConfig{Environment: cfg.Env, LogLevel: logger.DEBUG},
		os.Stdout,
	)
	l.Info("config loaded", "config", cfg)

	cliRequest, err := grpc.NewClient(
		cfg.Services.RequestServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer cliRequest.Close()

	cliUser, err := grpc.NewClient(
		cfg.Services.UserServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer cliUser.Close()

	requestRepo := grpcrequest.New(requestv1.NewProxyRequestServiceClient(cliRequest))
	userRepo := grpcuser.New(nil)

	application := app.New(l, &cfg.Proxy, proxysvc.New(requestRepo, userRepo))

	go func() {
		application.MustRun()
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	application.Stop()
}
