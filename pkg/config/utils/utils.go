package configutils

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Loading config from the environment variable.
func LoadConfigFromEnv[T any]() (*T, error) {
	var cfg T
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't load config from env %w", err)
	}

	return &cfg, nil
}

// Loading config from the file.
func LoadConfigFromFile[T any](configPath string) (*T, error) {
	// check presence of the file
	if _, err := os.Stat(configPath); err != nil {
		return nil, fmt.Errorf("couldn't open config file: %w", err)
	}

	var cfg T
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't read config from file: %w", err)
	}

	return &cfg, nil
}

// GetEnv returns value of environment variable ENV or local if ENV.
func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	return env
}
