package adapter

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/blink-io/x/kratos/v2/transport"
)

type Adapter interface {
	transport.Endpointer

	transport.Server

	Handler() http.Handler

	Kind() transport.Kind

	Listener() Listener
}

type Initializer interface {
	Init(context.Context, Options)
}

type Options struct {
	Network  string
	Address  string
	Endpoint *url.URL
	TlsConf  *tls.Config
	Handler  http.Handler
}
