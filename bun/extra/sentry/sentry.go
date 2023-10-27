package sentry

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/uptrace/bun"
)

type hook struct {
	hub *sentry.Hub
}

func New(options ...Option) bun.QueryHook {
	h := &hook{}
	for _, o := range options {
		o(h)
	}
	if h.hub == nil {
		h.hub = sentry.CurrentHub()
	}

	return h
}

func (q *hook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	if err := event.Err; err != nil {
		q.hub.CaptureException(err)
	}
	return ctx
}

func (q *hook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	if err := event.Err; err != nil {
		q.hub.CaptureException(err)
	}
}
