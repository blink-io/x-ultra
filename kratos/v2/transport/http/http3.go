package http

import (
	"crypto/tls"
	"net/http"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func IsHTTP3Transport(trans http.RoundTripper) bool {
	_, ok := trans.(*http3.RoundTripper)
	return ok
}

func RoundTripper(tlsConf *tls.Config) *http3.RoundTripper {
	qconf := new(quic.Config)
	return RoundTripperConf(tlsConf, qconf)
}

func RoundTripperConf(tlsConf *tls.Config, qconf *quic.Config) *http3.RoundTripper {
	rt := &http3.RoundTripper{
		TLSClientConfig: tlsConf,
		QuicConfig:      qconf,
	}
	return rt
}

func WithHTTP3RoundTripper(tlsConf *tls.Config) ClientOption {
	return WithTransport(RoundTripper(tlsConf))
}
