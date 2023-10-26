package bunlogging

import (
	"context"

	"github.com/uptrace/bun"
)

type hook struct {
}

func New() bun.QueryHook {
	return &hook{}
}

func (q *hook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (q *hook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
}

type Func func(string, ...any)

func (f Func) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	f(event.Query, event.QueryArgs...)
	return ctx
}

func (f Func) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
}

var _ bun.QueryHook = (Func)(nil)
