package i18n

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// grpcLoader loads by GRPC services
type grpcLoader struct {
	client    I18NClient
	endpoint  string
	languages []string
}

func NewGRPCLoader(cc grpc.ClientConnInterface, languages []string) Loader {
	client := NewI18NClient(cc)
	return &grpcLoader{client: client, languages: languages}
}

func (l *grpcLoader) Load(b Bundler) error {
	req := &ListLanguagesRequest{
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
			log("[WARN] unable to load message from gRPC: %s, endpoint: %s, reason: %s", v.Path, l.endpoint, verr.Error())
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

type grpcServer struct {
	UnimplementedI18NServer
	h EntryHandler
}

func newGrpcServer(h EntryHandler) *grpcServer {
	gsrv := &grpcServer{h: h}
	return gsrv
}

func (s *grpcServer) ListLanguages(ctx context.Context, req *ListLanguagesRequest) (*ListLanguagesResponse, error) {
	langs := req.Languages

	entries := make(map[string]*LanguageEntry)
	if s.h != nil {
		em := s.h.Handle(ctx, langs)
		for _, l := range langs {
			le := &LanguageEntry{
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
	res := &ListLanguagesResponse{
		Entries:   entries,
		Timestamp: ts,
	}

	return res, nil
}

func RegisterEntryHandler(gsrv *grpc.Server, h EntryHandler) {
	ss := newGrpcServer(h)
	RegisterI18NServer(gsrv, ss)
}

func RegisterEntryHandlerFunc(gsrv *grpc.Server, f EntryHandlerFunc) {
	ss := newGrpcServer(f)
	RegisterI18NServer(gsrv, ss)
}
