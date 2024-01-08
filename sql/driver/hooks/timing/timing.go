package timing

import (
	"context"
	"log"
	"time"

	"github.com/blink-io/x/sql/driver/hooks"
)

type ctxKey struct{}

type hook struct {
	logf func(string, ...any)
}

var _ hooks.Hooks = (*hook)(nil)

func New(ops ...Option) hooks.Hooks {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	if h.logf == nil {
		h.logf = log.Printf
	}
	return h
}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	newCtx := context.WithValue(ctx, ctxKey{}, time.Now())
	return newCtx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	if before, ok := ctx.Value(ctxKey{}).(time.Time); ok {
		h.logf("[SQLHooks] Executed SQL, timing cost [%s] for: %s", time.Since(before), query)
	}
	return ctx, nil
}
