package manager

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	managercfg "github.com/sazonovItas/go-simple-proxy/internal/config/manager"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
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
	_ = cfg

	logger := InitLogger("dev", os.Stdout)

	cli, err := client.NewClientWithOpts(
		client.WithVersion(cfg.DockerClient.ApiVersion),
		client.WithHost(cfg.DockerClient.Host),
		client.WithTimeout(cfg.DockerClient.Timeout),
	)
	if err != nil {
		logger.Error("failed to init client", slogger.Err(err), "config", *cfg)
		return
	}

	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        "go-proxy",
			ExposedPorts: nat.PortSet{"8123": struct{}{}},
			Env: []string{
				"ENV=dev", "PROXY_ID=test_proxy_id", "HOST=0.0.0.0", "PORT=8123",
				"READ_HEADER_TIMEOUT=1s", "READ_TIMEOUT=60s", "WRITE_TIMEOUT=60s", "SHUTDOWN_TIMEOUT=10s",
			},
		},
		&container.HostConfig{
			AutoRemove: true,
			PortBindings: nat.PortMap{"8123": []nat.PortBinding{
				{HostIP: "127.0.0.1", HostPort: "8123"},
			}},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"test-proxy",
	)
	if err != nil {
		log.Fatalf("error create container: %s", err.Error())
	}
	log.Printf("container created with id: %s", resp.ID)

	if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("error start container: %s", err.Error())
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()
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
