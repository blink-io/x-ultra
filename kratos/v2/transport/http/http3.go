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

// HTTP3Transport defines HTTP3 Transport RoundTripper
func HTTP3Transport(tlsConf *tls.Config) *http3.RoundTripper {
	qconf := new(quic.Config)
	return HTTP3TransportConf(tlsConf, qconf)
}

// HTTP3TransportConf defines HTTP3 Transport RoundTripper with custom quic.Config
func HTTP3TransportConf(tlsConf *tls.Config, qconf *quic.Config) *http3.RoundTripper {
	rt := &http3.RoundTripper{
		TLSClientConfig: tlsConf,
		QuicConfig:      qconf,
	}
	return rt
}

func WithHTTP3Transport(tlsConf *tls.Config) ClientOption {
	return WithTransport(HTTP3Transport(tlsConf))
}
