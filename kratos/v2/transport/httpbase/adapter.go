package httpbase

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/go-kratos/kratos/v2/transport"
)

type ServerAdapter interface {
	Handler() http.Handler

	Endpoint() (*url.URL, error)

	Start(context.Context) error

	Stop(context.Context) error

	Kind() transport.Kind

	Listener() Listener
}

type AdapterInit interface {
	Init(context.Context, *AdapterOptions)
}

type Validator interface {
	Validate(context.Context) error
}

type AdapterOptions struct {
	network  string
	address  string
	endpoint *url.URL
	tlsConf  *tls.Config
	handler  http.Handler
}

func (o *AdapterOptions) Network() string {
	return o.network
}

func (o *AdapterOptions) Address() string {
	return o.address
}

func (o *AdapterOptions) Endpoint() *url.URL {
	return o.endpoint
}

func (o *AdapterOptions) Handler() http.Handler {
	return o.handler
}

func (o *AdapterOptions) TLSConfig() *tls.Config {
	return o.tlsConf
}

type AdapterOption func(o *AdapterOptions)

func ApplyAdapterOptions(ops ...AdapterOption) *AdapterOptions {
	opts := new(AdapterOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func AdaptNetwork(network string) AdapterOption {
	return func(o *AdapterOptions) {
		o.network = network
	}
}

func AdaptAddress(address string) AdapterOption {
	return func(o *AdapterOptions) {
		o.address = address
	}
}

func AdaptTLSConfig(tlsConf *tls.Config) AdapterOption {
	return func(o *AdapterOptions) {
		o.tlsConf = tlsConf
	}
}

func AdaptHandler(handler http.Handler) AdapterOption {
	return func(o *AdapterOptions) {
		o.handler = handler
	}
}
