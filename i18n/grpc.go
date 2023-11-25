package i18n

import (
	"context"

	"google.golang.org/grpc"
)

// grpcLoader loads by GRPC services
type grpcLoader struct {
	client    I18NClient
	endpoint  string
	languages []string
}

func NewGRPCLoader(endpoint string, languages []string) (Loader, error) {
	cc, err := grpc.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	c := NewI18NClient(cc)
	return &grpcLoader{client: c, endpoint: endpoint}, nil
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
		if _, verr := b.LoadMessageFileBytes(v.Payload, v.Path); verr != nil {
			log("[WARN] unable to load message from gRPC: %s, endpoint: %s", v.Path, l.endpoint)
		}
	}
	return nil
}
