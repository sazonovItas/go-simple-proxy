package proxy

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"github.com/sazonovItas/go-simple-proxy/internal/proxy/models"
	proxyutils "github.com/sazonovItas/go-simple-proxy/internal/proxy/utils"
)

type RequestResponseRepo interface {
	Store(ctx context.Context, requestResponse *models.RequestResponse) error
}

type proxyRoundTripper struct {
	next   http.RoundTripper
	logger *slog.Logger
	repo   RequestResponseRepo
}

var _ http.RoundTripper = &proxyRoundTripper{}

func NewProxyRoundTripper(l *slog.Logger, repo RequestResponseRepo) *proxyRoundTripper {
	return &proxyRoundTripper{
		next:   http.DefaultTransport,
		logger: l,
		repo:   repo,
	}
}

func (prt *proxyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	dumpRequest, err := httputil.DumpRequest(r, true)
	if err != nil {
		prt.logger.Debug("error while dump request", "error", err.Error())
	} else {
		prt.logger.Debug("request dump", "request", fmt.Sprintf("%s %s HTTP/%d.%d", r.Method, r.URL.Path, r.ProtoMajor, r.ProtoMinor))
	}

	reqDump := proxyutils.ParseRequest(r)
	reqDump.ReceivedAt = time.Now()
	if r.Body != nil {
		var reader io.ReadCloser
		switch r.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(r.Body)
			if err != nil {
				prt.logger.Debug("error gzip.NewReader", "error", err.Error())
			}
			defer reader.Close()
		default:
			reader = r.Body
		}

		rawBody, err := io.ReadAll(reader)
		if err != nil {
			prt.logger.Debug("error while reading request body", "error", err.Error())
		}
		reqDump.Body = string(rawBody)

	}
	reqDump.Raw = string(dumpRequest)

	resp, err := prt.next.RoundTrip(r)
	if err != nil {
		return resp, err
	}

	dumpResponse, err := httputil.DumpResponse(resp, true)
	if err != nil {
		prt.logger.Debug("error while dump response", "error", err.Error())
	} else {
		prt.logger.Debug("response dump", "response", fmt.Sprintf("HTTP/%d.%d %d %s", resp.ProtoMajor, resp.ProtoMinor, resp.StatusCode, resp.Status))
	}

	respDump := proxyutils.ParseResponse(resp)
	if resp.Body != nil {
		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(resp.Body)
			if err != nil {
				prt.logger.Debug("error gzip.NewReader", "error", err.Error())
			}
			reader.Close()
		default:
			reader = resp.Body
		}

		rawBody, err := io.ReadAll(reader)
		if err != nil {
			prt.logger.Debug("error while reading response body", "error", err.Error())
		}

		if err := resp.Body.Close(); err != nil {
			prt.logger.Debug("error while closing response body", "error", err.Error())
		}

		body := io.NopCloser(bytes.NewReader(rawBody))
		resp.Body = body
		resp.ContentLength = int64(len(rawBody))
		resp.Header.Set("Content-Length", strconv.Itoa(len(rawBody)))
		respDump.Body = string(rawBody)
	}
	respDump.Raw = string(dumpResponse)
	respDump.ReceivedAt = time.Now()

	reqResp := models.RequestResponse{
		ProcessedAt: time.Now(),
		Request:     *reqDump,
		Response:    *respDump,
	}

	prt.logger.Info("processed request", "payload", reqResp)

	return resp, err
}
