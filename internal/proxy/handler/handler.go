package proxy

import (
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	proxyutils "github.com/sazonovItas/go-simple-proxy/internal/proxy/utils"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

type ProxyHandler struct {
	logger *slog.Logger
}

func NewProxyHandler(logger *slog.Logger) *ProxyHandler {
	return &ProxyHandler{
		logger: logger,
	}
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ph.logger.Debug(
		"new http request",
		"method", r.Method,
		"url", r.URL.String(),
		"headers", r.Header,
		"cookies", r.Cookies(),
		"remote_address", r.RemoteAddr,
	)

	if r.Header.Get("Proxy-Authorization") == "" {
		w.Header().Set("Proxy-Authenticate", "Basic realm=proxy")
		w.WriteHeader(http.StatusProxyAuthRequired)
		return
	}

	if r.Method == http.MethodConnect {
		ph.handleHTTPS(w, r)
		return
	}

	ph.handleHTTP(w, r)
}

func (ph *ProxyHandler) handleHTTPS(w http.ResponseWriter, r *http.Request) {
	ph.logger.Debug("hijacking connection", "src", r.RemoteAddr, "dest", r.URL.Host)
	clientConn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		ph.logger.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	ph.logger.Debug("connecting", "address", r.URL.Host)
	targetConn, err := net.Dial("tcp", r.URL.Host)
	if err != nil {
		ph.logger.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.WriteRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	ph.WriteRawResponse(clientConn, http.StatusOK, r)

	ph.logger.Debug("transferring", "from", r.RemoteAddr, "to", r.URL.Host)
	go func() {
		_, err = io.Copy(targetConn, clientConn)
		targetConn.Close()
	}()

	_, err = io.Copy(clientConn, targetConn)
	ph.logger.Debug("done transferring", "from", r.RemoteAddr, "to", r.URL.Host)
}

func (ph *ProxyHandler) WriteRawResponse(conn net.Conn, statusCode int, r *http.Request) {
	if err := proxyutils.WriteRawResponse(conn, statusCode, r); err != nil {
		ph.logger.Error("writing response", "error", err.Error())
	}
}

func (ph *ProxyHandler) handleHTTP(w http.ResponseWriter, r *http.Request) {
	ph.logger.Debug("hijacking connection", "src", r.RemoteAddr, "dest", r.URL.Host)
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.logger.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetHost := r.URL.Host
	if len(strings.Split(targetHost, ":")) == 1 {
		targetHost += ":80"
	}

	ph.logger.Debug("connecting", "address", r.URL.Host)
	targetConn, err := net.Dial("tcp", targetHost)
	if err != nil {
		ph.logger.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.WriteRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	clientDumpReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		ph.logger.Error("failed get dump request", "error", err.Error())
		ph.WriteRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	_, err = targetConn.Write(clientDumpReq)
	if err != nil {
		ph.logger.Error("failed write client request", "error", err.Error())
		ph.WriteRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	ph.logger.Debug("transferring", "from", r.RemoteAddr, "to", r.URL.Host)
	go func() {
		_, _ = io.Copy(targetConn, clientConn)
		targetConn.Close()
	}()

	_, _ = io.Copy(clientConn, targetConn)
	ph.logger.Debug("done transferring", "from", r.RemoteAddr, "to", r.URL.Host)
}
