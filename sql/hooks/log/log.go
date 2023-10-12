package log

import (
	"context"
	"log"

	"github.com/blink-io/x/sql/hooks"
)

type hook struct {
	f Fn
}

func New(o *Options) hooks.Hooks {
	if o == nil {
		o = DefaultOptions
	}
	f := o.Fn
	if f == nil {
		f = log.Printf
	}
	return &hook{f: f}
}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	h.f(query, args...)
	return ctx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}
