package thrift

import (
	"context"
	"net/http"
	"time"

	"github.com/blink-io/x/i18n"

	"github.com/apache/thrift/lib/go/thrift"
)

type loader struct {
	client    *I18NClient
	languages []string
	endpoint  string
}

func NewLoader(addr string, languages []string, ops ...Option) (i18n.Loader, error) {
	opt := applyOptions(ops...)
	framed := opt.framed
	useHTTP := opt.useHTTP
	protocolType := opt.protocol

	var cfg = &thrift.TConfiguration{
		ConnectTimeout: i18n.DefaultTimeout,
		SocketTimeout:  i18n.DefaultTimeout,
	}
	var transport thrift.TTransport
	var err error
	var headers = make(map[string]string)
	if useHTTP {
		transport, err = thrift.NewTHttpClientWithOptions(addr, thrift.THttpClientOptions{
			Client: &http.Client{
				Timeout: i18n.DefaultTimeout,
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
	case Compact:
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
		break
	case SimpleJSON:
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
		break
	case JSON:
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case Binary:
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
		break
	default:
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
		break
	}
	inproto := protocolFactory.GetProtocol(transport)
	outproto := protocolFactory.GetProtocol(transport)
	client := NewI18NClient(thrift.NewTStandardClient(inproto, outproto))
	ld := &loader{client: client, languages: languages}
	return ld, nil
}

func (l *loader) Load(b i18n.Bundler) error {
	req := NewListLanguagesRequest()
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
			i18n.GetLogger()("[WARN] unable to load message from Thrift service: %s, endpoint: %s, reason: %s", v.Path, l.endpoint, verr.Error())
		}
	}
	return nil
}

func LoadFromThrift(addr string, languages []string, ops ...Option) error {
	ld, err := NewLoader(addr, languages, ops...)
	if err != nil {
		return err
	}
	return ld.Load(i18n.Default())
}

var _ I18N = (*Handler)(nil)

type Handler struct {
	h i18n.EntryHandler
}

func NewHandler(h i18n.EntryHandler) *Handler {
	return &Handler{h: h}
}

func (s *Handler) ListLanguages(ctx context.Context, req *ListLanguagesRequest) (*ListLanguagesResponse, error) {
	langs := req.Languages

	entries := make(map[string]*LanguageEntry)
	if s.h != nil {
		em := s.h.Handle(ctx, langs)
		for _, l := range langs {
			le := NewLanguageEntry()
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

	res := NewListLanguagesResponse()
	res.Timestamp = time.Now().Unix()
	res.Entries = entries
	return res, nil
}

func NewTBinaryServer(addr string, eh i18n.EntryHandler, ops ...Option) (*thrift.TSimpleServer, error) {
	opt := applyOptions(ops...)

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
		ConnectTimeout: i18n.DefaultTimeout,
		SocketTimeout:  i18n.DefaultTimeout,
	}
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(cfg)

	processor := NewI18NProcessor(NewHandler(eh))
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	return server, nil
}
