package slogger

import (
	"io"
	"log/slog"

	sldiscard "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl/handlers/discard"
	slpretty "github.com/sazonovItas/go-simple-proxy/pkg/logger/sl/handlers/pretty"
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
