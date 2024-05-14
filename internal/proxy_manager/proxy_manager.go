package proxymanager

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const configPathEnv = "CONFIG_PATH"

func Run() {
	// cfg, err := configutils.LoadCfgFromFile[proxymanagercfg.Config](os.Getenv(configPathEnv))
	// if err != nil {
	// 	log.Fatalf("faild to load proxy manager config: %s", err.Error())
	// }
	// _ = cfg

	cli, err := client.NewClientWithOpts(client.WithVersion("1.44"))
	if err != nil {
		log.Fatalf("error create client from env: %s", err.Error())
	}

	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        "go-proxy",
			ExposedPorts: nat.PortSet{"8123": struct{}{}},
			Env: []string{
				"ENV=dev", "PROXY_ID=test_proxy_id", "ADDRESS=0.0.0.0:8123",
				"READ_HEADER_TIMEOUT=1s", "IDLE_TIMEOUT=60s", "SHUTDOWN_TIMEOUT=10s",
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
