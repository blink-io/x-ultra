package httpbase

import (
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func IsHTTP3Transport(trans http.RoundTripper) bool {
	_, ok := trans.(*http3.RoundTripper)
	return ok
}
