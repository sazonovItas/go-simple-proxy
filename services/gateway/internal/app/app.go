package app

import (
	"log/slog"

	accountv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/account/v1"
	authv1 "github.com/sazonovItas/proxy-manager/services/auth/pkg/pb/auth/v1"
	managerv1 "github.com/sazonovItas/proxy-manager/services/proxy-manager/pkg/pb/manager/v1"
	requestv1 "github.com/sazonovItas/proxy-manager/services/proxy-request/pkg/pb/request/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcmanagerapi "github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter/grpcapi/manager"
	grpcrequestapi "github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter/grpcapi/request"
	grpcuserapi "github.com/sazonovItas/proxy-manager/services/gateway/internal/adapter/grpcapi/user"
	httpapp "github.com/sazonovItas/proxy-manager/services/gateway/internal/app/http"
	"github.com/sazonovItas/proxy-manager/services/gateway/internal/config"
)

type App struct {
	HTTPServer *httpapp.App

	cliUser    *grpc.ClientConn
	cliProxy   *grpc.ClientConn
	cliRequest *grpc.ClientConn
}

func New(l *slog.Logger, cfg *config.Config) *App {
	cliUser, err := grpc.NewClient(
		cfg.Services.AuthSvcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	cliProxy, err := grpc.NewClient(
		cfg.Services.ManagerSvcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	cliRequest, err := grpc.NewClient(
		cfg.Services.RequestSvcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	userApi := grpcuserapi.New(authv1.NewAuthClient(cliUser), accountv1.NewAccountClient(cliUser))
	requestApi := grpcrequestapi.New(requestv1.NewProxyRequestServiceClient(cliRequest))
	proxyApi := grpcmanagerapi.New(managerv1.NewProxyManagerClient(cliProxy))

	httpServer := httpapp.New(l, &cfg.HTTPServer, userApi, requestApi, proxyApi)

	return &App{
		HTTPServer: httpServer,

		cliUser:    cliUser,
		cliProxy:   cliProxy,
		cliRequest: cliRequest,
	}
}

func (a *App) Stop() {
}
