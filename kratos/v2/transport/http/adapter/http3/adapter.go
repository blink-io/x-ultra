package http3

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"

	"github.com/blink-io/x/kratos/v2/internal/endpoint"
	"github.com/blink-io/x/kratos/v2/internal/host"
	xtransport "github.com/blink-io/x/kratos/v2/transport"
	xadapter "github.com/blink-io/x/kratos/v2/transport/http/adapter"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

type server = http3.Server

type Options = xadapter.Options

type ExtraOption func(*adapter)

func Listener(ln http3.QUICEarlyListener) ExtraOption {
	return func(a *adapter) {
		a.ln = ln
	}
}

func QConfig(qconf *quic.Config) ExtraOption {
	return func(a *adapter) {
		a.qconf = qconf
	}
}

type adapter struct {
	srv      *server
	network  string
	address  string
	tlsConf  *tls.Config
	endpoint *url.URL
	ln       http3.QUICEarlyListener
	qconf    *quic.Config
}

var DefaultOptions = Options{
	Network: "udp",
	Address: ":0",
}

func NewDefault() xadapter.Adapter {
	return NewAdapter(DefaultOptions)
}

func NewAdapter(opts Options, eops ...ExtraOption) xadapter.Adapter {
	a := new(adapter)
	for _, o := range eops {
		o(a)
	}
	a.Init(context.Background(), opts)
	return a
}

func (s *adapter) Init(ctx context.Context, opts Options) {
	s.network = opts.Network
	s.address = opts.Address
	s.tlsConf = opts.TLSConf
	s.endpoint = opts.Endpoint
	s.qconf = new(quic.Config)
	s.srv = &http3.Server{
		Addr:       s.address,
		TLSConfig:  s.tlsConf,
		QuicConfig: s.qconf,
		Handler:    opts.Handler,
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
	return xtransport.KindHTTP3
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

func (s *adapter) Listener() xadapter.Listener {
	return s.ln
}
