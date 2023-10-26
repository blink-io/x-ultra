package sentry

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
)

type hook struct {
	hub *sentry.Hub
}

var _ rueidishook.Hook = (*hook)(nil)

func New(ops ...Option) (rueidishook.Hook, error) {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	return h, nil
}

func (h *hook) Do(c rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	r := c.Do(ctx, cmd)
	h.handleSingleError(r)
	return r
}

func (h *hook) DoMulti(c rueidis.Client, ctx context.Context, multi ...rueidis.Completed) []rueidis.RedisResult {
	ra := c.DoMulti(ctx, multi...)
	for _, r := range ra {
		if err := r.Error(); err != nil {
			h.handleSingleError(r)
		}
	}
	return ra
}

func (h *hook) DoCache(c rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) rueidis.RedisResult {
	r := c.DoCache(ctx, cmd, ttl)
	h.handleSingleError(r)
	return r
}

func (h *hook) DoMultiCache(c rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) []rueidis.RedisResult {
	return c.DoMultiCache(ctx, multi...)
}

func (h *hook) Receive(c rueidis.Client, ctx context.Context, sub rueidis.Completed, fn func(msg rueidis.PubSubMessage)) error {
	if err := c.Receive(ctx, sub, fn); err != nil {
		h.hub.CaptureException(err)
		return err
	}
	return nil
}

func (h *hook) handleSingleError(r rueidis.RedisResult) {
	if err := r.Error(); err != nil {
		h.hub.CaptureException(err)
	}
}
