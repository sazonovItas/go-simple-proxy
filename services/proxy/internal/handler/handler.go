package proxyhandler

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

type requestUsecase interface {
	Save(ctx context.Context, r *entity.Request) error
}

type ProxyHandler struct {
	proxyID string
	l       *slog.Logger

	requestUsc requestUsecase
}

func NewProxyHandler(proxyId string, logger *slog.Logger, requestUsc requestUsecase) *ProxyHandler {
	return &ProxyHandler{
		l:          logger,
		proxyID:    proxyId,
		requestUsc: requestUsc,
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
	createdAt := time.Now()

	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.l.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetConn, err := net.Dial("tcp", r.URL.Host)
	if err != nil {
		ph.l.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	ph.writeRawResponse(clientConn, http.StatusOK, r)

	upload, download := ph.transfering(clientConn, targetConn)

	request := entity.Request{
		ProxyID:       ph.proxyID,
		ProxyName:     ph.proxyID,
		ProxyUserID:   "itas",
		ProxyUserIP:   r.RemoteAddr,
		ProxyUserName: "itas",
		Host:          strings.Split(r.URL.Host, ":")[0],
		Upload:        upload,
		Download:      download,
		CreatedAt:     createdAt,
	}
	ph.l.Info("content length", "request", request)

	if err := ph.requestUsc.Save(context.Background(), &request); err != nil {
		ph.l.Error("failed to save request", slogger.Err(err))
		return
	}
}

func (ph *ProxyHandler) handleHTTP(w http.ResponseWriter, r *http.Request) {
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.l.Error("hijack failed", "error", err.Error())
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
		ph.l.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	clientDumpReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		ph.l.Error("failed get dump request", "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	_, err = targetConn.Write(clientDumpReq)
	if err != nil {
		ph.l.Error("failed write client request", "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	upload, download := ph.transfering(clientConn, targetConn)
	_, _ = upload, download
}

func (ph *ProxyHandler) transfering(clientConn, targetConn net.Conn) (upload, download int64) {
	quitch := make(chan struct{})
	defer close(quitch)

	go func() {
		upload, _ = io.Copy(targetConn, clientConn)
		targetConn.Close()
		quitch <- struct{}{}
	}()

	download, _ = io.Copy(clientConn, targetConn)
	<-quitch

	return
}

func (ph *ProxyHandler) writeRawResponse(conn net.Conn, statusCode int, r *http.Request) {
	if err := WriteRawResponse(conn, statusCode, r); err != nil {
		ph.l.Error("writing response", "error", err.Error())
	}
}
