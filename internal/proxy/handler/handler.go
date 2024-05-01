package proxy

import (
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"slices"
	"strings"

	configproxy "github.com/sazonovItas/go-simple-proxy/internal/config/proxy"
	"github.com/sazonovItas/go-simple-proxy/internal/proxy/common"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

type ProxyHandler struct {
	logger      *slog.Logger
	certManager *CertManager
	blockList   []string
}

func NewProxyHandler(
	logger *slog.Logger,
	certCfg configproxy.ProxySecrets,
	blockList []string,
) *ProxyHandler {
	certManager, err := NewCertManager(certCfg)
	if err != nil {
		panic(err)
	}

	return &ProxyHandler{
		logger:      logger,
		certManager: certManager,
		blockList:   blockList,
	}
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ph.logger.Info("new request", "remote_addr", r.RemoteAddr)

	ph.logger.Debug(
		"new http request",
		"method", r.Method,
		"url", r.URL.String(),
		"headers", r.Header,
		"cookies", r.Cookies(),
	)

	host := strings.Split(r.URL.Host, ":")[0]
	if slices.Index(ph.blockList, host) != -1 {
		ph.logger.Info("access forbidden", "host", host)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method == http.MethodConnect {
		ph.handleHTTPS(w, r)
		return
	}

	ph.handleHTTP(w, r)
}

func (ph *ProxyHandler) handleHTTP(w http.ResponseWriter, r *http.Request) {
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

	ph.logger.Debug("transfering", "from", r.RemoteAddr, "to", r.URL.Host)
	go func() {
		_, _ = io.Copy(targetConn, clientConn)
		targetConn.Close()
	}()

	_, _ = io.Copy(clientConn, targetConn)
	ph.logger.Debug("done transfering", "from", r.RemoteAddr, "to", r.URL.Host)
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

	ph.logger.Debug("transfering", "from", r.RemoteAddr, "to", r.URL.Host)
	go func() {
		_, err = io.Copy(targetConn, clientConn)
		targetConn.Close()
	}()

	_, err = io.Copy(clientConn, targetConn)
	ph.logger.Debug("done transfering", "from", r.RemoteAddr, "to", r.URL.Host)
}

func (ph *ProxyHandler) WriteRawResponse(conn net.Conn, statusCode int, r *http.Request) {
	if err := common.WriteRawResponse(conn, statusCode, r); err != nil {
		ph.logger.Error("writing response", "error", err.Error())
	}
}
