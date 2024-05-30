package proxyhandler

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

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
		ph.log.Error("writing response", "error", err.Error())
	}
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
