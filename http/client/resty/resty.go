package resty

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func HTTP3Client(tlsConf *tls.Config) *resty.Client {
	c := resty.New()
	c.SetTransport(&http3.RoundTripper{
		TLSClientConfig: tlsConf,
		QuicConfig:      new(quic.Config),
	})
	return c
}
