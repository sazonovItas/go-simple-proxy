package proxyutils

import (
	"fmt"
	"net"
	"net/http"

	"github.com/sazonovItas/go-simple-proxy/internal/proxy/models"
)

func ParseResponse(r *http.Response) *models.Response {
	resp := &models.Response{
		StatusCode: r.StatusCode,
	}

	headers := make(http.Header)
	for k, values := range r.Header {
		headers[k] = append(headers[k], values...)
	}
	resp.Headers = headers

	cookies := make(map[string]string)
	for _, v := range r.Cookies() {
		cookies[v.Name] = v.Value
	}
	resp.Cookies = cookies

	return resp
}

func WriteRawResponse(conn net.Conn, statusCode int, r *http.Request) error {
	_, err := fmt.Fprintf(
		conn,
		"HTTP/%d.%d %03d %s\r\n\r\n",
		r.ProtoMajor,
		r.ProtoMinor,
		statusCode,
		http.StatusText(statusCode),
	)
	return err
}
