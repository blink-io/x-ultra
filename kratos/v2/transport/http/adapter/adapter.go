package adapter

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/go-kratos/kratos/v2/transport"
)

type Adapter interface {
	transport.Endpointer

	Handler() http.Handler

	Start(context.Context) error

	Stop(context.Context) error

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
	TLSConf  *tls.Config
	Handler  http.Handler
}