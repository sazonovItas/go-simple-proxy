package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger(l *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		l.Info("logger middleware is using")

		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			next.ServeHTTP(w, r)
			t2 := time.Since(t1)

			defer l.Info("request completed",
				slog.String("request_id", RequestIdFromContext(r.Context())),
				slog.String("method", r.Method),
				slog.String("host", r.URL.Host),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("duration", t2.String()),
			)
		}

		return http.HandlerFunc(fn)
	}
}

func NewLoggerResponseWriter(w http.ResponseWriter) *loggerResponseWriter {
	return &loggerResponseWriter{ResponseWriter: w}
}

type loggerResponseWriter struct {
	http.ResponseWriter
}
