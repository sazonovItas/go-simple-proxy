package zaplogger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewSugaredLogger(level zapcore.Level, w io.Writer) *zap.Logger {
	var cfg zapcore.EncoderConfig = zap.NewProductionEncoderConfig()

	encoder := zapcore.NewJSONEncoder(cfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(w), level)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}
