package sentry

import (
	"context"

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

func (h *hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (h *hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	if err := cmd.Err(); err != nil {
		h.hub.CaptureException(err)
	}
	return nil
}

func (h *hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (h *hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil {
			h.hub.CaptureException(err)
		}
	}
	return nil
}

func (h *hook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (h *hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return next
}

func (h *hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return next
}
