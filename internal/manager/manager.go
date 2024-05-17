package manager

import (
	"context"
	"errors"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	managercfg "github.com/sazonovItas/go-simple-proxy/internal/config/manager"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
	"github.com/sazonovItas/go-simple-proxy/internal/manager/engine"
	slogger "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl"
)

const (
	configPathEnv = "CONFIG_PATH"
	local         = "local"
	development   = "dev"
	production    = "prod"
)

func Run() {
	cfg, err := configutils.LoadCfgFromFile[managercfg.Config](os.Getenv(configPathEnv))
	if err != nil {
		log.Fatalf("faild to load proxy manager config: %s", err.Error())
	}

	logger := InitLogger("dev", os.Stdout)
	logger.Info("init config", "config", *cfg)

	// cli, err := client.NewClientWithOpts(
	// 	client.WithVersion(cfg.DockerClient.ApiVersion),
	// 	client.WithHost(cfg.DockerClient.Host),
	// 	client.WithTimeout(cfg.DockerClient.Timeout),
	// )
	// if err != nil {
	// 	logger.Error("failed to init client", slogger.Err(err), "config", *cfg)
	// 	return
	// }
	//
	// resp, err := cli.ContainerCreate(
	// 	context.Background(),
	// 	&container.Config{
	// 		Image:        "go-proxy",
	// 		ExposedPorts: nat.PortSet{"8123": struct{}{}},
	// 		Env: []string{
	// 			"ENV=dev", "PROXY_ID=test_proxy_id", "PORT=8123",
	// 			"READ_HEADER_TIMEOUT=1s", "READ_TIMEOUT=60s", "WRITE_TIMEOUT=60s", "SHUTDOWN_TIMEOUT=10s",
	// 		},
	// 	},
	// 	&container.HostConfig{
	// 		AutoRemove: true,
	// 		PortBindings: nat.PortMap{"8123": []nat.PortBinding{
	// 			{HostIP: "127.0.0.1", HostPort: "8123"},
	// 		}},
	// 	},
	// 	&network.NetworkingConfig{},
	// 	&v1.Platform{},
	// 	"test-proxy",
	// )
	// if err != nil {
	// 	log.Fatalf("error create container: %s", err.Error())
	// }
	// log.Printf("container created with id: %s", resp.ID)
	//
	// if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
	// 	log.Fatalf("error start container: %s", err.Error())
	// }

	engine, err := engine.NewEngine(cfg.ProxyManager.ProxyImage.Image, engine.DockerClientConfig{
		ApiVersion: cfg.DockerClient.ApiVersion,
		Timeout:    cfg.DockerClient.Timeout,
		Host:       cfg.DockerClient.Host,
	})
	if err != nil {
		logger.Error("failed init engine", slogger.Err(err))
		return
	}

	err = engine.Run(cfg.ProxyManager.Proxies)
	if err != nil {
		logger.Error("failed run proxy manager", slogger.Err(err))
		return
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()

	// TODO: change shutdown timeout with config value
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer func() {
		cancel()

		if shutdownCtx.Err() != nil && !errors.Is(shutdownCtx.Err(), context.Canceled) {
			logger.Warn("proxy manager shutdown with error", slogger.Err(shutdownCtx.Err()))
		}
	}()

	if err := engine.Shutdown(shutdownCtx); err != nil {
		logger.Error("engine shuted down with error", "error", err.Error())
	}

	logger.Info("proxy manager is shuted down")
}

func InitLogger(env string, out io.Writer) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case development:
		logger = slogger.NewPrettyLogger(slog.LevelInfo, out)
	case production:
		logger = slogger.NewPrettyLogger(slog.LevelWarn, out)
	default:
		logger = slogger.NewPrettyLogger(slog.LevelDebug, out)
	}

	return logger
}
