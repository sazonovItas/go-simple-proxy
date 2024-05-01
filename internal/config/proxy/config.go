package configproxy

type Proxy struct {
	Address   string       `yaml:"address"`
	Secrets   ProxySecrets `yaml:"secrets"`
	BlockList []string     `yaml:"block-list"`
}

type ProxySecrets struct {
	Key  string `yaml:"key"`
	Cert string `yaml:"cert"`
}
