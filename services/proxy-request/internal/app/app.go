package app

import (
	"log/slog"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"

	grpcapp "github.com/sazonovItas/proxy-manager/proxy-request/internal/app/grpc"
	httpapp "github.com/sazonovItas/proxy-manager/proxy-request/internal/app/http"
	"github.com/sazonovItas/proxy-manager/proxy-request/internal/config"
)

type App struct {
	GRPCServer grpcapp.App
	HTTPServer httpapp.App
}

func New(
	l *slog.Logger,
	db *postgresdb.DB,
	grpcCfg *config.GRPCServerConfig,
	httpCfg *config.HTTPServerConfig,
) *App {
	return &App{}
}
