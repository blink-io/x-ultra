package http3

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"

	"github.com/blink-io/x/kratos/v2/internal/endpoint"
	"github.com/blink-io/x/kratos/v2/internal/host"
	"github.com/blink-io/x/kratos/v2/transport/httpbase"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

type server = http3.Server

type Options = httpbase.AdapterOptions

type adapter struct {
	srv      *server
	network  string
	address  string
	tlsConf  *tls.Config
	endpoint *url.URL
	ln       http3.QUICEarlyListener
	qconf    *quic.Config
}

var DefaultOptions = newDefaultOptions()

func newDefaultOptions() *Options {
	opts := httpbase.ApplyAdapterOptions(
		httpbase.AdapterNetwork("udp"),
		httpbase.AdapterAddress(":0"),
	)
	return opts
}

func NewAdapter(opts *Options) httpbase.ServerAdapter {
	a := new(adapter)
	a.Init(context.Background(), opts)
	return a
}

func (s *adapter) Init(ctx context.Context, opts *Options) {
	s.network = opts.Network()
	s.address = opts.Address()
	s.tlsConf = opts.TLSConfig()
	s.endpoint = opts.Endpoint()
	s.qconf = new(quic.Config)
	s.srv = &http3.Server{
		Addr:       s.address,
		TLSConfig:  s.tlsConf,
		QuicConfig: s.qconf,
		Handler:    opts.Handler(),
	}
}

func (s *adapter) Validate(ctx context.Context) error {
	if s.srv.TLSConfig == nil {
		return errors.New("http3 adapter: tlsConf is required")
	}
	return nil
}

func (s *adapter) Handler() http.Handler {
	return s.srv.Handler
}

func (s *adapter) Kind() transport.Kind {
	return httpbase.KindHTTP3
}

// Start start the HTTP server.
func (s *adapter) Start(ctx context.Context) error {
	if err := s.listenAndEndpoint(); err != nil {
		return err
	}

	log.Infof("[HTTP3] server listening on: %s", s.ln.Addr().String())

	err := s.srv.ServeListener(s.ln)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop the HTTP server.
func (s *adapter) Stop(ctx context.Context) error {
	log.Info("[HTTP3] server stopping")
	return s.srv.Close()
}

func (s *adapter) listenAndEndpoint() error {
	if s.ln == nil {
		ln, err := quic.ListenAddrEarly(s.address, http3.ConfigureTLSConfig(s.tlsConf), s.qconf)
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
		s.endpoint = endpoint.NewEndpoint("https", addr)
	}
	return nil
}

func (s *adapter) Endpoint() (*url.URL, error) {
	if err := s.listenAndEndpoint(); err != nil {
		return nil, err
	}
	return s.endpoint, nil
}

func (s *adapter) Listener() httpbase.Listener {
	return s.ln
}
