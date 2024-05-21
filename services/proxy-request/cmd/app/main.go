package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	configutils "github.com/sazonovItas/proxy-manager/pkg/config/utils"
	"github.com/sazonovItas/proxy-manager/pkg/logger"
	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	pb_request "github.com/sazonovItas/proxy-manager/proxy-request/api/proto/pb"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/config"
	grpcrequest "github.com/sazonovItas/proxy-manager/proxy-request/internal/handler/grpc/request"
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

	handler := grpcrequest.NewRequestHandler(nil)

	grpcserver := grpc.NewServer(grpc.ConnectionTimeout(cfg.RPCServer.ReadTimeout))
	pb_request.RegisterProxyRequestServiceServer(grpcserver, handler)
	reflection.Register(grpcserver)

	listener, err := net.Listen("tcp", cfg.RPCServer.Address)
	if err != nil {
		l.Error("faild listen on address", "address", cfg.RPCServer.Address, slogger.Err(err))
		return
	}
	defer listener.Close()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		l.Info("GRPC server started", "address", cfg.RPCServer.Address)
		if err := grpcserver.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			l.Error("server shutdown", slogger.Err(err))
		}
	}()

	grpcclient, err := grpc.NewClient(
		cfg.RPCServer.Address,
		grpc.WithIdleTimeout(cfg.RPCServer.ReadTimeout),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err == nil {
		cli := pb_request.NewProxyRequestServiceClient(grpcclient)
		go func() {
			for {
				select {
				case <-ctx.Done():
					l.Info("client closed")
					return
				case <-time.After(time.Second * 2):
					msg, err := cli.SaveProxyRequest(context.Background(), &pb_request.SaveRequest{
						Request: &pb_request.ProxyRequest{
							Host: "itas",
						},
					})
					if err != nil {
						l.Error("grpc clietn error", slogger.Err(err))
						continue
					}

					l.Info("client response", "id", msg.Id)
				}
			}
		}()
	} else {
		l.Error("client connection", slogger.Err(err))
	}

	go func() {}()

	<-ctx.Done()

	grpcserver.GracefulStop()
}
