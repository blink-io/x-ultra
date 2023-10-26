package sentry

import (
	"context"
	"net"

	"github.com/getsentry/sentry-go"
	"github.com/redis/go-redis/v9"
)

type hook struct {
	hub *sentry.Hub
}

var _ redis.Hook = (*hook)(nil)

func New(ops ...Option) (redis.Hook, error) {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	return h, nil
}

func (h *hook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		cc, err := next(ctx, network, addr)
		if err != nil {
			h.hub.CaptureException(err)
		}
		return cc, err
	}
}

func (h *hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		err := next(ctx, cmd)
		if err != nil {
			h.hub.CaptureException(err)
		}
		return err
	}
}

func (h *hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		err := next(ctx, cmds)
		if err != nil {
			h.hub.CaptureException(err)
		}
		return err
	}
}
