package http

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/url"

	"github.com/blink-io/x/kratos/v2/internal/endpoint"
	"github.com/blink-io/x/kratos/v2/internal/host"
	xadapter "github.com/blink-io/x/kratos/v2/transport/http/adapter"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

type (
	server = http.Server

	Options = xadapter.Options

	ExtraOption func(*adapter)
)

func Listener(ln net.Listener) ExtraOption {
	return func(a *adapter) {
		a.ln = ln
	}
}

type adapter struct {
	srv      *server
	network  string
	address  string
	ln       net.Listener
	tlsConf  *tls.Config
	endpoint *url.URL
}

var DefaultOptions = Options{
	Network: "tcp",
	Address: ":0",
}

func NewDefault() xadapter.Adapter {
	return NewAdapter(DefaultOptions)
}

func NewAdapter(opts Options, eops ...ExtraOption) xadapter.Adapter {
	a := new(adapter)
	a.Init(context.Background(), opts)
	for _, o := range eops {
		o(a)
	}
	return a
}

func (s *adapter) Init(ctx context.Context, opts Options) {
	s.network = opts.Network
	s.address = opts.Address
	s.tlsConf = opts.TLSConf
	s.endpoint = opts.Endpoint
	s.srv = &http.Server{
		Addr:      s.address,
		TLSConfig: s.tlsConf,
		Handler:   opts.Handler,
	}
}

func (s *adapter) Validate(ctx context.Context) error {
	return nil
}

// Start start the HTTP server.
func (s *adapter) Start(ctx context.Context) error {
	if err := s.listenAndEndpoint(); err != nil {
		return err
	}
	s.srv.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	log.Infof("[HTTP] server listening on: %s", s.ln.Addr().String())
	var err error
	if s.tlsConf != nil {
		err = s.srv.ServeTLS(s.ln, "", "")
	} else {
		err = s.srv.Serve(s.ln)
	}
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop the HTTP server.
func (s *adapter) Stop(ctx context.Context) error {
	log.Info("[HTTP] server stopping")
	return s.srv.Shutdown(ctx)
}

func (s *adapter) listenAndEndpoint() error {
	if s.ln == nil {
		ln, err := net.Listen(s.network, s.address)
		if err != nil {
			return err
		}
		s.ln = ln
	}
	if s.endpoint == nil {
		addr, err := host.Extract(s.address, s.ln)
		if err != nil {
			return err
		}
		s.endpoint = endpoint.NewEndpoint(endpoint.Scheme("http", s.tlsConf != nil), addr)
	}
	return nil
}

func (s *adapter) Endpoint() (*url.URL, error) {
	if err := s.listenAndEndpoint(); err != nil {
		return nil, err
	}
	return s.endpoint, nil
}

func (s *adapter) Handler() http.Handler {
	return s.srv.Handler
}

func (s *adapter) Kind() transport.Kind {
	return transport.KindHTTP
}

func (s *adapter) Listener() xadapter.Listener {
	return s.ln
}
