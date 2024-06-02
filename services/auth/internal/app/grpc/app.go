package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sazonovItas/proxy-manager/services/auth/internal/config"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/handler/grpc/account"
	"github.com/sazonovItas/proxy-manager/services/auth/internal/handler/grpc/auth"
)

type App struct {
	log *slog.Logger

	cfg        *config.GRPCServerConfig
	grpcServer *grpc.Server
	authSvc    auth.AuthService
}

func New(
	cfg *config.GRPCServerConfig,
	l *slog.Logger,
	authSvc auth.AuthService,
	accountSvc account.AccountService,
) *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			// logging.StartCall, logging.FinishCall,
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			l.Error("recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(l), loggingOpts...),
	))

	auth.Register(gRPCServer, auth.New(authSvc))
	account.Register(gRPCServer, account.New(accountSvc))

	return &App{
		log: l,

		cfg:        cfg,
		grpcServer: gRPCServer,
		authSvc:    authSvc,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			l.Log(ctx, slog.Level(lvl), msg, fields...)
		},
	)
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.With(slog.String("op", op)).
		Info("starting grpc server", slog.String("address", l.Addr().String()))

	if err := a.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping grpc server")

	a.grpcServer.GracefulStop()
}
