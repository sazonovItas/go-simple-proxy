package middleware

import (
	"log/slog"
	"net/http"
)

func Panic(l *slog.Logger) func(next http.Handler) http.Handler {
	const op = "handler.middleware.Panic"

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := l.With(slog.String("panic middleware", op))

			defer func() {
				if r := recover(); r != nil {
					log.Error("recovered from panic", "info", r)
				}
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
