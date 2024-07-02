//go:build go1.20

package shared

import (
	"bufio"
	"net"
	"net/http"
)

func (sw *SessionResponseWriter) Flush() {
	http.NewResponseController(sw.ResponseWriter).Flush()
}

func (sw *SessionResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return http.NewResponseController(sw.ResponseWriter).Hijack()
}
