package logger

import (
	"io"
	"log/slog"

	"go.uber.org/zap/zapcore"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"
)

type Level string

type LogConfig struct {
	Environment string
	LogLevel    Level
}

const (
	TRACE Level = "TRACE"
	DEBUG Level = "DEBUG"
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	PANIC Level = "PANIC"

	local       string = "local"
	development string = "dev"
	production  string = "prod"
)

type LoggerConfig struct {
	Env   string
	Level Level
}

func NewSlogLogger(cfg LogConfig, out io.Writer) *slog.Logger {
	switch cfg.Environment {
	case production:
		return slogger.NewDiscardLogger(logLevelToSlog(cfg.LogLevel), out)
	default:
		return slogger.NewPrettyLogger(logLevelToSlog(INFO), out)
	}
}

func logLevelToSlog(level Level) (slLevel slog.Level) {
	switch level {
	case TRACE:
		slLevel = slog.LevelDebug
	case DEBUG:
		slLevel = slog.LevelDebug
	case INFO:
		slLevel = slog.LevelDebug
	case WARN:
		slLevel = slog.LevelWarn
	case ERROR:
		slLevel = slog.LevelError
	case PANIC:
		slLevel = slog.LevelError
	}

	return
}

func logLevelToZap(level Level) (zapLevel zapcore.Level) {
	switch level {
	case TRACE:
		zapLevel = zapcore.DebugLevel
	case DEBUG:
		zapLevel = zapcore.DebugLevel
	case INFO:
		zapLevel = zapcore.InfoLevel
	case WARN:
		zapLevel = zapcore.WarnLevel
	case ERROR:
		zapLevel = zapcore.ErrorLevel
	case PANIC:
		zapLevel = zapcore.PanicLevel
	}

	return
}
