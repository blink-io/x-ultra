package i18n

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	i18nthrift "github.com/blink-io/x/i18n/thrift"
)

type thriftOptions struct {
	protocol   i18nthrift.Protocol
	useHTTP    bool
	framed     bool
	buffered   bool
	bufferSize int
	tlsConfig  *tls.Config
}

func applyTOptions(ops ...ThriftOption) *thriftOptions {
	opt := new(thriftOptions)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

type ThriftOption func(*thriftOptions)

func WithTProtocol(protocol i18nthrift.Protocol) ThriftOption {
	return func(o *thriftOptions) {
		o.protocol = protocol
	}
}

func WithTFramed(framed bool) ThriftOption {
	return func(o *thriftOptions) {
		o.framed = framed
	}
}

func WithUseHTTP() ThriftOption {
	return func(o *thriftOptions) {
		o.useHTTP = true
	}
}

func WithBuffered(bufferSize int) ThriftOption {
	return func(o *thriftOptions) {
		o.buffered = true
		o.bufferSize = bufferSize
	}
}

func WithTLSConfig(tlsConfig *tls.Config) ThriftOption {
	return func(o *thriftOptions) {
		o.tlsConfig = tlsConfig
	}
}

type thriftLoader struct {
	client    *i18nthrift.I18NClient
	languages []string
	endpoint  string
}

func NewThriftLoader(addr string, languages []string, ops ...ThriftOption) (Loader, error) {
	opt := applyTOptions(ops...)
	framed := opt.framed
	useHTTP := opt.useHTTP
	protocolType := opt.protocol

	var cfg = &thrift.TConfiguration{
		ConnectTimeout: DefaultTimeout,
		SocketTimeout:  DefaultTimeout,
	}
	var transport thrift.TTransport
	var err error
	var headers = make(map[string]string)
	if useHTTP {
		transport, err = thrift.NewTHttpClientWithOptions(addr, thrift.THttpClientOptions{
			Client: &http.Client{
				Timeout: DefaultTimeout,
			},
		})
		if len(headers) > 0 {
			httptrans := transport.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
		if framed {
			transport = thrift.NewTFramedTransportConf(transport, cfg)
		}
	}

	if err != nil {
		return nil, err
	}

	var transportFactory thrift.TTransportFactory
	if opt.buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(opt.bufferSize)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	if vtrans, verr := transportFactory.GetTransport(transport); verr != nil {
		return nil, verr
	} else {
		transport = vtrans
	}

	// Open transport
	if err := transport.Open(); err != nil {
		return nil, err
	}

	var protocolFactory thrift.TProtocolFactory
	switch protocolType {
	case i18nthrift.Compact:
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
		break
	case i18nthrift.SimpleJSON:
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
		break
	case i18nthrift.JSON:
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case i18nthrift.Binary:
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
		break
	default:
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
		break
	}
	inproto := protocolFactory.GetProtocol(transport)
	outproto := protocolFactory.GetProtocol(transport)
	client := i18nthrift.NewI18NClient(thrift.NewTStandardClient(inproto, outproto))
	ld := &thriftLoader{client: client, languages: languages}
	return ld, nil
}

func (l *thriftLoader) Load(b Bundler) error {
	req := i18nthrift.NewListLanguagesRequest()
	req.Languages = l.languages

	res, err := l.client.ListLanguages(context.Background(), req)
	if err != nil {
		return err
	}
	for _, v := range res.Entries {
		// Ignore invalid
		if !v.Valid {
			continue
		}
		if _, verr := b.LoadMessageFileBytes(v.Payload, v.Path); verr != nil {
			log("[WARN] unable to load message from Thrift service: %s, endpoint: %s, reason: %s", v.Path, l.endpoint, verr.Error())
		}
	}
	return nil
}

func (b *Bundle) LoadFromThrift(addr string, languages []string, ops ...ThriftOption) error {
	ld, err := NewThriftLoader(addr, languages, ops...)
	if err != nil {
		return err
	}
	return ld.Load(b)
}

func LoadFromThrift(addr string, languages []string, ops ...ThriftOption) error {
	ld, err := NewThriftLoader(addr, languages, ops...)
	if err != nil {
		return err
	}
	return ld.Load(bb)
}

var _ i18nthrift.I18N = (*ThriftHandler)(nil)

type ThriftHandler struct {
	h EntryHandler
}

func NewThriftHandler(h EntryHandler) *ThriftHandler {
	return &ThriftHandler{h: h}
}

func (s *ThriftHandler) ListLanguages(ctx context.Context, req *i18nthrift.ListLanguagesRequest) (*i18nthrift.ListLanguagesResponse, error) {
	langs := req.Languages

	entries := make(map[string]*i18nthrift.LanguageEntry)
	if s.h != nil {
		em := s.h.Handle(ctx, langs)
		for _, l := range langs {
			le := i18nthrift.NewLanguageEntry()
			le.Language = l
			if e := em[l]; e != nil {
				le.Path = e.Path
				le.Payload = e.Payload
				le.Valid = true
			} else {
				le.Valid = false
			}
			entries[l] = le
		}
	}

	res := i18nthrift.NewListLanguagesResponse()
	res.Timestamp = time.Now().Unix()
	res.Entries = entries
	return res, nil
}

func NewTBinaryServer(addr string, h EntryHandler, ops ...ThriftOption) (*thrift.TSimpleServer, error) {
	opt := applyTOptions(ops...)

	var err error
	var serverTransport thrift.TServerTransport
	if opt.tlsConfig != nil {
		serverTransport, err = thrift.NewTSSLServerSocket(addr, opt.tlsConfig)
	} else {
		serverTransport, err = thrift.NewTServerSocket(addr)
	}
	if err != nil {
		return nil, err
	}

	cfg := &thrift.TConfiguration{
		ConnectTimeout: DefaultTimeout,
		SocketTimeout:  DefaultTimeout,
	}
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(cfg)

	processor := i18nthrift.NewI18NProcessor(NewThriftHandler(h))
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	server.SetLogger(func(msg string) {
		slog.Default().Info(msg)
	})
	return server, nil
}
