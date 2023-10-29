package sentry

import (
	"context"

	"github.com/blink-io/x/sql/hooks"
	"github.com/getsentry/sentry-go"
)

type hook struct {
	hub *sentry.Hub
}

var _ hooks.Hooks = (*hook)(nil)
var _ hooks.OnErrorer = (*hook)(nil)

func New(ops ...Option) (hooks.Hooks, error) {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	return h, nil
}

func (h *hook) OnError(ctx context.Context, err error, query string, args ...interface{}) error {
	if err != nil {
		h.hub.CaptureException(err)
	}
	return err
}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}
