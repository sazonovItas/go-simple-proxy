package logger

import "log/slog"

type (
	Level      int8
	EnvSetting string
)

const (
	ERROR Level = 8
	WARN  Level = 4
	INFO  Level = 0
	DEBUG Level = -4
	TRACE Level = 4

	LOCAL       EnvSetting = "local"
	DEVELOPMENT EnvSetting = "development"
	PRODUCTION  EnvSetting = "production"
)

type LoggerConfig struct {
	Env   string
	Level Level
}

func NewSlogLogger() *slog.Logger {
	panic("need to implement")
}
