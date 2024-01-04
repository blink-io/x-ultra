package requests

import (
	"crypto/tls"

	"github.com/carlmjohnson/requests"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

type Builder = requests.Builder

func B() *Builder {
	return requests.New()
}

func HTTP3(tlsConf *tls.Config) *Builder {
	return HTTP3Conf(tlsConf, new(quic.Config))
}

func HTTP3Conf(tlsConf *tls.Config, qconf *quic.Config) *Builder {
	cc := requests.New().
		Transport(&http3.RoundTripper{
			TLSClientConfig: tlsConf,
			QuicConfig:      qconf,
		})
	return cc
}
