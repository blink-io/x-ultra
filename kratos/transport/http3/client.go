package http3

import (
	"crypto/tls"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func RoundTripper(tlsConf *tls.Config) *http3.RoundTripper {
	qconf := new(quic.Config)
	rt := &http3.RoundTripper{TLSClientConfig: tlsConf, QuicConfig: qconf}
	return rt
}
