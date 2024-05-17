package sldiscard

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"
	"strings"
	"time"
)

type DiscardLogFormat struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Message string    `json:"msg"`
	Payload string    `json:"payload,omitempty"`
}

type DiscardHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type DiscardHandler struct {
	slog.Handler

	l     *stdLog.Logger
	attrs []slog.Attr
}

func (opts *DiscardHandlerOptions) NewDiscardHandler(out io.Writer) *DiscardHandler {
	return &DiscardHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}
}

func (h *DiscardHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *DiscardHandler) Handle(ctx context.Context, r slog.Record) error {
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var (
		b   []byte
		err error
	)
	if len(fields) > 0 {
		b, err = json.Marshal(fields)
		if err != nil {
			return err
		}
	}

	level := strings.ToLower(r.Level.String())
	msg := DiscardLogFormat{
		Time:    r.Time,
		Level:   level,
		Message: r.Message,
		Payload: string(b),
	}

	logMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	h.l.Println(string(logMsg))

	return nil
}

func (h *DiscardHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &DiscardHandler{
		Handler: h.Handler.WithAttrs(attrs),
		l:       h.l,
		attrs:   attrs,
	}
}

// TODO: implement.
func (h *DiscardHandler) WithGroup(name string) slog.Handler {
	return &DiscardHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
