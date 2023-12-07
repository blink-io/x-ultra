package logging

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

type hook struct {
	logf func(string, ...any)
}

func New(ops ...Option) bun.QueryHook {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	if h.logf == nil {
		h.logf = log.Printf
	}
	return h
}

func (q *hook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (q *hook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
}

// Func defines
type Func func(format string, args ...any)

func (f Func) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	f("[QueryHook] Executed SQL, query: %s, args: %q", event.Query, event.QueryArgs)
	return ctx
}

func (f Func) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
}

type CtxFunc func(ctx context.Context, format string, args ...any)

func (f CtxFunc) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	f(ctx, "[QueryHook] Executed SQL, query: %s, args: %q", event.Query, event.QueryArgs)
	return ctx
}

func (f CtxFunc) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
}
