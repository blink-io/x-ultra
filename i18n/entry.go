package i18n

import "context"

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
