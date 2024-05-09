package proxy

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"
	"strings"
	"time"

	proxyutils "github.com/sazonovItas/go-simple-proxy/internal/proxy/utils"
)

const (
	HTTP  = "http"
	HTTPS = "https"
)

type ProxyHandler struct {
	logger    *slog.Logger
	rt        http.RoundTripper
	blockList []string
}

func NewProxyHandler(
	logger *slog.Logger,
	blockList []string,
) *ProxyHandler {
	return &ProxyHandler{
		logger:    logger,
		rt:        NewProxyRoundTripper(logger, nil),
		blockList: blockList,
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

	// if _, err := r.Cookie("ProxyAuth"); err != nil && errors.Is(err, http.ErrNoCookie) {
	// 	http.Redirect(w, r, "http://google.com", http.StatusTemporaryRedirect)
	// 	return
	// }

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

func (ph *ProxyHandler) handleHTTPS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Proxy-Authorization") == "" {
		w.Header().Set("Proxy-Authenticate", "Digest realm=\""+r.Host+"\"")
		w.WriteHeader(http.StatusProxyAuthRequired)
		return
	}

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

	cookie := http.Cookie{
		Name:     "ProxyAuth",
		Value:    "hi",
		Expires:  time.Now().Add(time.Second * 3600),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	ph.logger.Info("cookie", "cookie", cookie.String())

	_, err = fmt.Fprintf(
		clientConn,
		"HTTP/%d.%d %03d %s\r\nSet-Cookie: %s\r\n\r\n",
		r.ProtoMajor,
		r.ProtoMinor,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		cookie.String(),
	)

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

func handleSingle(client *http.Client, inReq *http.Request) ([]byte, error) {
	ctx := inReq.Context()
	outReq := inReq.Clone(ctx)

	outReq.RequestURI = ""
	if inReq.ContentLength == 0 {
		outReq.Body = nil
	}

	if outReq.Body != nil {
		defer outReq.Body.Close()
	}

	if outReq.Header == nil {
		outReq.Header = make(http.Header)
	}
	outReq.Close = false

	if _, ok := outReq.Header["User-Agent"]; !ok {
		outReq.Header.Set("User-Agent", "")
	}

	resp, err := client.Do(outReq)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	return httputil.DumpResponse(resp, true)
}

func changeRequestToTarget(req *http.Request, targetHost string, proto string) error {
	if proto != HTTPS {
		return nil
	}

	if !strings.HasPrefix(targetHost, "https") {
		targetHost = "https://" + targetHost
	}

	targetUrl, err := url.Parse(targetHost)
	if err != nil {
		return err
	}

	targetUrl.Path = req.URL.Path
	targetUrl.RawQuery = req.URL.RawQuery
	req.URL = targetUrl

	req.RequestURI = ""
	return nil
}
