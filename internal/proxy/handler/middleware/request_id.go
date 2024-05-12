package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// RequestIdKey is key for request id in request context
const requestIdKey key = "request_id"

// RequestId is middleware for request id.
// Request id go along with context value (string) under the key RequestIdKey
func RequestId() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), requestIdKey, generateId())
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func RequestIdFromContext(ctx context.Context) string {
	if requestId, ok := ctx.Value(requestIdKey).(string); ok {
		return requestId
	}

	return ""
}

func generateId() string {
	return uuid.New().String()
}
