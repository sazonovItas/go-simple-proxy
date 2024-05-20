package config

type Config struct {
	HTTPServer HTTPServerConfig
	RPCServer  RPCServerConfig
}

type HTTPServerConfig struct{}

type RPCServerConfig struct{}
