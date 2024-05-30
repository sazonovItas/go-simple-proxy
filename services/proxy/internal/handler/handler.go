package proxyhandler

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	slogger "github.com/sazonovItas/proxy-manager/pkg/logger/sl"

	"github.com/sazonovItas/proxy-manager/services/proxy/internal/entity"
	"github.com/sazonovItas/proxy-manager/services/proxy/internal/handler/middleware"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

type ProxyService interface {
	Save(ctx context.Context, r *entity.Request) error
	Login(ctx context.Context, login, password string) (entity.Token, error)
}

type ProxyHandler struct {
	id      string
	timeout time.Duration

	log      *slog.Logger
	proxySvc ProxyService
}

func New(
	id string,
	timeout time.Duration,

	logger *slog.Logger,
	proxySvc ProxyService,
) *ProxyHandler {
	return &ProxyHandler{
		id:      id,
		timeout: timeout,

		log:      logger,
		proxySvc: proxySvc,
	}
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	creds, _ := middleware.ProxyUserCreditanalsFromContext(r.Context())

	token, err := ph.proxySvc.Login(r.Context(), creds.Username, creds.Password)
	if err != nil {
		middleware.ProxyBasicAuthFailed(w, "")
		return
	}

	if r.Method == http.MethodConnect {
		ph.handleHTTPS(w, r, token)
		return
	}

	if len(strings.Split(r.URL.Host, ":")) == 1 {
		r.URL.Host += ":80"
	}
	ph.handleHTTP(w, r, token)
}

func (ph *ProxyHandler) handleHTTPS(
	w http.ResponseWriter,
	r *http.Request,
	token entity.Token,
) {
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.log.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetConn, err := net.DialTimeout("tcp", r.URL.Host, ph.timeout)
	if err != nil {
		ph.log.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	ph.writeRawResponse(clientConn, http.StatusOK, r)

	createdAt := time.Now()
	upload, download := ph.transfering(clientConn, targetConn)
	request := entity.Request{
		UserID:    token.UserID,
		ProxyID:   ph.id,
		RemoteIP:  r.RemoteAddr,
		Host:      strings.Split(r.URL.Host, ":")[0],
		Upload:    upload,
		Download:  download,
		CreatedAt: createdAt,
	}

	ph.SaveRequest(context.Background(), &request)
}

func (ph *ProxyHandler) handleHTTP(w http.ResponseWriter, r *http.Request, token entity.Token) {
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.log.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetConn, err := net.DialTimeout("tcp", r.URL.Host, ph.timeout)
	if err != nil {
		ph.log.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	clientDumpReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		ph.log.Error("failed get dump request", "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	_, err = targetConn.Write(clientDumpReq)
	if err != nil {
		ph.log.Error("failed write client request", "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}

	createdAt := time.Now()
	upload, download := ph.transfering(clientConn, targetConn)
	request := entity.Request{
		UserID:    token.UserID,
		ProxyID:   ph.id,
		RemoteIP:  r.RemoteAddr,
		Host:      strings.Split(r.URL.Host, ":")[0],
		Upload:    upload,
		Download:  download,
		CreatedAt: createdAt,
	}

	ph.SaveRequest(context.Background(), &request)
}

func (ph *ProxyHandler) SaveRequest(ctx context.Context, r *entity.Request) {
	if err := ph.proxySvc.Save(ctx, r); err != nil {
		ph.log.Error("failed to save request", slogger.Err(err))
		return
	}
}
