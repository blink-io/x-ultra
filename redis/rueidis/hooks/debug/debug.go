package debug

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
)

type hook struct {
	logf func(string, ...any)
}

var _ rueidishook.Hook = (*hook)(nil)

func New(ops ...Option) rueidishook.Hook {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	if h.logf == nil {
		h.logf = log.Printf
	}
	return h
}

func (h *hook) Do(c rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	h.logf("Redis CMD: [%s]", cmdstr(cmd.Commands()))
	return c.Do(ctx, cmd)
}

func (h *hook) DoMulti(c rueidis.Client, ctx context.Context, multi ...rueidis.Completed) []rueidis.RedisResult {
	for _, m := range multi {
		h.logf("Redis CMD: [%s]", cmdstr(m.Commands()))
	}
	rr := c.DoMulti(ctx, multi...)
	return rr
}

func (h *hook) DoCache(c rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) rueidis.RedisResult {
	h.logf("Redis CMD: [%s]", cmdstr(cmd.Commands()))
	r := c.DoCache(ctx, cmd, ttl)
	return r
}

func (h *hook) DoMultiCache(c rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) []rueidis.RedisResult {
	for _, m := range multi {
		h.logf("Redis CMD: [%s]", cmdstr(m.Cmd.Commands()))
	}
	rr := c.DoMultiCache(ctx, multi...)
	return rr
}

func (h *hook) Receive(c rueidis.Client, ctx context.Context, sub rueidis.Completed, fn func(msg rueidis.PubSubMessage)) error {
	h.logf("Redis CMD: [%s]", cmdstr(sub.Commands()))
	err := c.Receive(ctx, sub, fn)
	return err
}

func cmdstr(cmds []string) string {
	return strings.Join(cmds, " ")
}
