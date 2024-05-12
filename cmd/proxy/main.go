package main

import (
	"log"

	configproxy "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
	configutils "github.com/sazonovItas/go-simple-proxy/internal/config/utils"
	"github.com/sazonovItas/go-simple-proxy/internal/proxy"
)

func main() {
	cfg, err := configutils.LoadCfgFromFile[configproxy.Config](
		"./config/" + configutils.GetEnv() + "/proxy.yml",
	)
	if err != nil {
		log.Fatalf("failed load config with error: %s", err.Error())
		return
	}

	proxy.Run(cfg)
}
