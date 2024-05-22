package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pgrequest "github.com/sazonovItas/proxy-manager/proxy-request/internal/adapter/pgrepo/request"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/config"
	grpcrequest "github.com/sazonovItas/proxy-manager/proxy-request/internal/handler/grpc/request"
	requestsvc "github.com/sazonovItas/proxy-manager/proxy-request/internal/service/request"
	pb_request "github.com/sazonovItas/proxy-manager/proxy-request/pkg/pb"
)

func main() {
	cfg, err := configutils.LoadConfigFromEnv[config.Config]()
	if err != nil {
		log.Fatalf("failed load config from env: %s", err.Error())
	}

	l := logger.NewSlogLogger(
		logger.LogConfig{Environment: cfg.Env, LogLevel: logger.DEBUG},
		os.Stdout,
	)
	l.Info("init config", "config", cfg)

	// connect to db
	// TODO: move to internal init
	db, err := postgresdb.Connect(
		context.Background(),
		cfg.Storage.Uri,
		&postgresdb.ConnectionOptions{},
	)
	if err != nil {
		l.Error("failed connect to database", slogger.Err(err))
		return
	}
	defer db.Close()

	// init request repo
	// TODO: move to internal init
	requestRepo := pgrequest.NewRequestRepository("proxy_requests", db)
	_ = requestsvc.NewRequestService(requestRepo)

	// init grpc handler
	// TODO: move to internal init
	handler := grpcrequest.NewRequestHandler(l, requestRepo)
	grpcserver := grpc.NewServer(grpc.ConnectionTimeout(cfg.RPCServer.Timeout))
	pb_request.RegisterProxyRequestServiceServer(grpcserver, handler)
	reflection.Register(grpcserver)

	listener, err := net.Listen("tcp", cfg.RPCServer.Address)
	if err != nil {
		l.Error("faild listen on address", "address", cfg.RPCServer.Address, slogger.Err(err))
		return
	}
	defer listener.Close()

	// graceful shutdown
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		l.Info("RPC server started", "address", cfg.RPCServer.Address)
		if err := grpcserver.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			l.Error("failed serve connection", slogger.Err(err))
		}
	}()
	<-ctx.Done()

	grpcserver.GracefulStop()
	l.Info("server stopped")
}
