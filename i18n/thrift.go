package i18n

import (
	"context"
	"net/http"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	i18nthrift "github.com/blink-io/x/i18n/thrift"
)

type TOptions struct {
	protocol   i18nthrift.Protocol
	useHTTP    bool
	framed     bool
	buffered   bool
	bufferSize int
}

func applyTOptions(ops ...TOption) *TOptions {
	opt := new(TOptions)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

type TOption func(*TOptions)

func WithTProtocol(protocol i18nthrift.Protocol) TOption {
	return func(o *TOptions) {
		o.protocol = protocol
	}
}

func WithTFramed(framed bool) TOption {
	return func(o *TOptions) {
		o.framed = framed
	}
}

func WithUseHTTP() TOption {
	return func(o *TOptions) {
		o.useHTTP = true
	}
}

func WithBuffered(bufferSize int) TOption {
	return func(o *TOptions) {
		o.buffered = true
		o.bufferSize = bufferSize
	}
}

type thriftLoader struct {
	client    *i18nthrift.I18NClient
	languages []string
	endpoint  string
}

func NewThriftLoader(addr string, languages []string, ops ...TOption) (Loader, error) {
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
	req := &i18nthrift.ListLanguagesRequest{
		Languages: l.languages,
	}
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

func (b *Bundle) LoadFromThrift(addr string, languages []string, ops ...TOption) error {
	ld, err := NewThriftLoader(addr, languages, ops...)
	if err != nil {
		return err
	}
	return ld.Load(b)
}

func LoadFromThrift(addr string, languages []string, ops ...TOption) error {
	ld, err := NewThriftLoader(addr, languages, ops...)
	if err != nil {
		return err
	}
	return ld.Load(bb)
}

var _ i18nthrift.I18N = (*thriftHandler)(nil)

type thriftHandler struct {
	h EntryHandler
}

func (s *thriftHandler) ListLanguages(ctx context.Context, req *i18nthrift.ListLanguagesRequest) (*i18nthrift.ListLanguagesResponse, error) {
	langs := req.Languages

	entries := make(map[string]*i18nthrift.LanguageEntry)
	if s.h != nil {
		em := s.h.Handle(ctx, langs)
		for _, l := range langs {
			le := &i18nthrift.LanguageEntry{
				Language: l,
			}
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

	ts := time.Now().Unix()
	res := &i18nthrift.ListLanguagesResponse{
		Timestamp: ts,
		Entries:   entries,
	}
	return res, nil
}

func NewTBinaryServer(addr string, h EntryHandler) (*thrift.TSimpleServer, error) {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return nil, err
	}

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)

	processor := i18nthrift.NewI18NProcessor(&thriftHandler{h: h})
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return server, nil
}
