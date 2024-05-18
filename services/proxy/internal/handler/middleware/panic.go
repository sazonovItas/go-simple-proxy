package middleware

import (
	"log/slog"
	"net/http"
)

func Panic(l *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					l.Error("recovered from panic", "info", r)
				}
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
