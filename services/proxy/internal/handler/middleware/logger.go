package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Logger(l *slog.Logger) func(next http.Handler) http.Handler {
	const op = "handler.middleware.Logger"

	return func(next http.Handler) http.Handler {
		log := l.With("logger middleware", op)

		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			next.ServeHTTP(w, r)
			t2 := time.Since(t1)

			defer log.Info("request completed",
				slog.String("protocol", fmt.Sprintf("HTTP/%d.%d", r.ProtoMajor, r.ProtoMinor)),
				slog.String("method", r.Method),
				slog.String("host", r.URL.Host),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("duration", t2.String()),
			)
		}

		return http.HandlerFunc(fn)
	}
}
