package proxyhandler

import (
	"fmt"
	"net"
	"net/http"
)

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
