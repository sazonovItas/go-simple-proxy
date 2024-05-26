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

type proxyService interface {
	Save(ctx context.Context, r *entity.Request) error
	Login(ctx context.Context, username, passwordHash string) (*entity.User, error)
}

type ProxyHandler struct {
	id      string
	name    string
	timeout time.Duration

	l        *slog.Logger
	proxySvc proxyService
}

func New(
	id, name string,
	timeout time.Duration,

	logger *slog.Logger,
	proxySvc proxyService,
) *ProxyHandler {
	return &ProxyHandler{
		id:      id,
		name:    name,
		timeout: timeout,

		l:        logger,
		proxySvc: proxySvc,
	}
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	creds, _ := middleware.ProxyUserCreditanalsFromContext(r.Context())
	user, err := ph.proxySvc.Login(r.Context(), creds.Username, creds.Password)
	if err != nil {
		middleware.ProxyBasicAuthFailed(w, "")
		return
	}

	if r.Method == http.MethodConnect {
		ph.handleHTTPS(w, r, user)
		return
	}

	if len(strings.Split(r.URL.Host, ":")) == 1 {
		r.URL.Host += ":80"
	}
	ph.handleHTTP(w, r, user)
}

func (ph *ProxyHandler) handleHTTPS(
	w http.ResponseWriter,
	r *http.Request,
	userCreds *entity.User,
) {
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.l.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetConn, err := net.DialTimeout("tcp", r.URL.Host, ph.timeout)
	if err != nil {
		ph.l.Error("connection failed", "address", r.URL.Host, "error", err.Error())
		ph.writeRawResponse(clientConn, http.StatusInternalServerError, r)
		return
	}
	defer targetConn.Close()

	ph.writeRawResponse(clientConn, http.StatusOK, r)

	createdAt := time.Now()
	upload, download := ph.transfering(clientConn, targetConn)
	request := entity.Request{
		ProxyID:       ph.id,
		ProxyName:     ph.name,
		ProxyUserID:   userCreds.ID,
		ProxyUserIP:   r.RemoteAddr,
		ProxyUserName: userCreds.Username,
		Host:          strings.Split(r.URL.Host, ":")[0],
		Upload:        upload,
		Download:      download,
		CreatedAt:     createdAt,
	}

	ph.SaveRequest(context.Background(), &request)
}

func (ph *ProxyHandler) handleHTTP(w http.ResponseWriter, r *http.Request, userCreds *entity.User) {
	rc := http.NewResponseController(w)
	_ = rc.EnableFullDuplex()

	clientConn, _, err := rc.Hijack()
	if err != nil {
		ph.l.Error("hijack failed", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	targetConn, err := net.DialTimeout("tcp", r.URL.Host, ph.timeout)
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

	createdAt := time.Now()

	upload, download := ph.transfering(clientConn, targetConn)
	request := entity.Request{
		ProxyID:       ph.id,
		ProxyName:     ph.name,
		ProxyUserID:   userCreds.ID,
		ProxyUserIP:   r.RemoteAddr,
		ProxyUserName: userCreds.Username,
		Host:          strings.Split(r.URL.Host, ":")[0],
		Upload:        upload,
		Download:      download,
		CreatedAt:     createdAt,
	}

	ph.SaveRequest(context.Background(), &request)
}

func (ph *ProxyHandler) SaveRequest(ctx context.Context, r *entity.Request) {
	if err := ph.proxySvc.Save(ctx, r); err != nil {
		ph.l.Error("failed to save request", slogger.Err(err))
		return
	}
}
