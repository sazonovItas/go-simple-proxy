package slogger

import (
	"io"
	"log/slog"

	sldiscard "github.com/sazonovItas/proxy-manager/pkg/logger/sl/handlers/discard"
	slpretty "github.com/sazonovItas/proxy-manager/pkg/logger/sl/handlers/pretty"
)

func NewPrettyLogger(level slog.Level, out io.Writer) *slog.Logger {
	opts := slpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewPrettyHandler(out)
	return slog.New(handler)
}

func NewDiscardLogger(level slog.Level, out io.Writer) *slog.Logger {
	opts := sldiscard.DiscardHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewDiscardHandler(out)
	return slog.New(handler)
}

func Err(err error) slog.Attr {
	if err == nil {
		slog.String("error", "nil")
	}

	return slog.String("error", err.Error())
}
