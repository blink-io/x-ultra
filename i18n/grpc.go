package i18n

import (
	"context"
	"time"

	i18ngrpc "github.com/blink-io/x/i18n/grpc"
	"google.golang.org/grpc"
)

// grpcLoader loads by GRPC services
type grpcLoader struct {
	client    i18ngrpc.I18NClient
	endpoint  string
	languages []string
}

func NewGRPCLoader(cc grpc.ClientConnInterface, languages []string) Loader {
	client := i18ngrpc.NewI18NClient(cc)
	return &grpcLoader{client: client, languages: languages}
}

func (b *Bundle) LoadFromGRPC(cc grpc.ClientConnInterface, languages []string) error {
	return NewGRPCLoader(cc, languages).Load(b)
}

func LoadFromGRPC(cc grpc.ClientConnInterface, languages []string) error {
	return NewGRPCLoader(cc, languages).Load(bb)
}

func (l *grpcLoader) Load(b Bundler) error {
	req := &i18ngrpc.ListLanguagesRequest{
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
			log("[WARN] unable to load message from gRPC service: %s, endpoint: %s, reason: %s", v.Path, l.endpoint, verr.Error())
		}
	}
	return nil
}

type (
	Entry struct {
		Path     string
		Language string
		Valid    bool
		Payload  []byte
	}

	EntryHandler interface {
		Handle(ctx context.Context, languages []string) map[string]*Entry
	}
)

type EntryHandlerFunc func(ctx context.Context, languages []string) map[string]*Entry

func (f EntryHandlerFunc) Handle(ctx context.Context, languages []string) map[string]*Entry {
	return f(ctx, languages)
}

type Entries map[string]*Entry

func (e Entries) Handle(ctx context.Context, languages []string) map[string]*Entry {
	ne := make(map[string]*Entry)
	if len(e) == 0 {
		return ne
	}
	for _, l := range languages {
		if ee, ok := e[l]; ok {
			ne[l] = ee
		}
	}
	return ne
}

type grpcServer struct {
	i18ngrpc.UnimplementedI18NServer
	h EntryHandler
}

func newGrpcServer(h EntryHandler) *grpcServer {
	gsrv := &grpcServer{h: h}
	return gsrv
}

func (s *grpcServer) ListLanguages(ctx context.Context, req *i18ngrpc.ListLanguagesRequest) (*i18ngrpc.ListLanguagesResponse, error) {
	langs := req.Languages

	entries := make(map[string]*i18ngrpc.LanguageEntry)
	if s.h != nil {
		em := s.h.Handle(ctx, langs)
		for _, l := range langs {
			le := &i18ngrpc.LanguageEntry{
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

	res := &i18ngrpc.ListLanguagesResponse{
		Entries:   entries,
		Timestamp: time.Now().Unix(),
	}

	return res, nil
}

func RegisterEntryHandler(gsrv *grpc.Server, h EntryHandler) {
	ss := newGrpcServer(h)
	i18ngrpc.RegisterI18NServer(gsrv, ss)
}

func RegisterEntryHandlerFunc(gsrv *grpc.Server, f EntryHandlerFunc) {
	ss := newGrpcServer(f)
	i18ngrpc.RegisterI18NServer(gsrv, ss)
}
