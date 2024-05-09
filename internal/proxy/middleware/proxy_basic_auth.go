package middleware

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func ProxyBasicAuth(l *slog.Logger, realm string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := getProxyBasicAuth(r)
			l.Info("auth from user", "username", user, "password", pass, "ok", ok)
			if !ok {
				proxyBasicAuthFailed(w, realm)
				return
			}

			l.Info("auth from user", "username", user, "password", pass)

			next.ServeHTTP(w, r)
		})
	}
}

func getProxyBasicAuth(r *http.Request) (string, string, bool) {
	proxyBasicAuth := r.Header.Get("Proxy-Authorization")
	if proxyBasicAuth == "" {
		return "", "", false
	}

	base64BasicAuth := strings.Split(proxyBasicAuth, " ")
	if len(base64BasicAuth) != 2 || base64BasicAuth[0] != "Basic" {
		return "", "", false
	}

	basicAuth, err := base64.RawStdEncoding.DecodeString(base64BasicAuth[1])
	if err != nil {
		return "", "", false
	}

	creditanals := strings.Split(string(basicAuth), ":")
	if len(creditanals) != 2 {
		return "", "", false
	}

	return creditanals[0], creditanals[1], true
}

func proxyBasicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("Proxy-Authenticate", fmt.Sprintf("Basic realm=%s", realm))
	w.WriteHeader(http.StatusProxyAuthRequired)
}
