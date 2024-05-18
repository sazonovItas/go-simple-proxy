package proxyhandler

import (
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	proxyutils "github.com/sazonovItas/proxy-manager/services/proxy/internal/utils"
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
	if r.Method == http.MethodConnect {
		ph.handleHTTPS(w, r)
		return
	}

	ph.handleHTTP(w, r)
}

func (ph *ProxyHandler) handleHTTPS(w http.ResponseWriter, r *http.Request) {
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.logger.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetConn, err := net.Dial("tcp", r.URL.Host)
	if err != nil {
		ph.logger.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	ph.writeRawResponse(clientConn, http.StatusOK, r)

	upload, download := ph.transfering(clientConn, targetConn)
	_, _ = upload, download
}

func (ph *ProxyHandler) handleHTTP(w http.ResponseWriter, r *http.Request) {
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

	targetConn, err := net.Dial("tcp", targetHost)
	if err != nil {
		ph.logger.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	clientDumpReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		ph.logger.Error("failed get dump request", "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	_, err = targetConn.Write(clientDumpReq)
	if err != nil {
		ph.logger.Error("failed write client request", "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	upload, download := ph.transfering(clientConn, targetConn)
	_, _ = upload, download
}

func (ph *ProxyHandler) transfering(clientConn, targetConn net.Conn) (upload, download int64) {
	go func() {
		upload, _ = io.Copy(targetConn, clientConn)
		targetConn.Close()
	}()

	download, _ = io.Copy(clientConn, targetConn)
	return
}

func (ph *ProxyHandler) writeRawResponse(conn net.Conn, statusCode int, r *http.Request) {
	if err := proxyutils.WriteRawResponse(conn, statusCode, r); err != nil {
		ph.logger.Error("writing response", "error", err.Error())
	}
}
