package proxymanager

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	proxymanagercfg "github.com/sazonovItas/go-simple-proxy/internal/config/proxy_manager"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
)

const configPathEnv = "CONFIG_PATH"

func Run() {
	cfg, err := configutils.LoadCfgFromFile[proxymanagercfg.Config](os.Getenv(configPathEnv))
	if err != nil {
		log.Fatalf("faild to load proxy manager config: %s", err.Error())
	}
	_ = cfg

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()
}
