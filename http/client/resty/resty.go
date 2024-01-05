package resty

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func HTTP3Client(tlsConf *tls.Config) *resty.Client {
	return HTTP3ClientConf(tlsConf, new(quic.Config))
}

func HTTP3ClientConf(tlsConf *tls.Config, qconf *quic.Config) *resty.Client {
	cc := resty.New()
	cc.SetTransport(&http3.RoundTripper{
		TLSClientConfig: tlsConf,
		QuicConfig:      qconf,
	})
	return cc
}
