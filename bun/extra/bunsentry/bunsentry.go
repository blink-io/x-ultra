package bunsentry

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/uptrace/bun"
)

const ctxKey = "bun-sentry"

type QueryHook struct {
	*sentry.Hub
}

var _ bun.QueryHook = (*QueryHook)(nil)

func NewQueryHook(options ...Option) *QueryHook {
	h := &QueryHook{}
	for _, o := range options {
		o(h)
	}
	if h.Hub == nil {
		h.Hub = sentry.CurrentHub()
	}
	return h
}

func (q *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	if err := event.Err; err != nil {
		sentry.CaptureException(err)
	}
	return ctx
}

func (q *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	if err := event.Err; err != nil {
		sentry.CaptureException(err)
	}
}
