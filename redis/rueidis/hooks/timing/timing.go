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
	before := time.Now()
	cstr := cmdstr(cmd.Commands())
	res := c.Do(ctx, cmd)
	cost := time.Since(before)
	h.logf("Timing cost [%s] for [Do]: [%s]", cost.String(), cstr)
	return res
}

func (h *hook) DoMulti(c rueidis.Client, ctx context.Context, multi ...rueidis.Completed) []rueidis.RedisResult {
	before := time.Now()
	for _, m := range multi {
		h.logf("Redis CMD: [%s]", cmdstr(m.Commands()))
	}
	rr := c.DoMulti(ctx, multi...)
	cost := time.Since(before)
	h.logf("Timing cost for [DoMulti]: [%s]", cost.String())
	return rr
}

func (h *hook) DoCache(c rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) rueidis.RedisResult {
	before := time.Now()
	cstr := cmdstr(cmd.Commands())
	r := c.DoCache(ctx, cmd, ttl)
	cost := time.Since(before)
	h.logf("Timing cost [%s] for [DoCache]: [%s]", cost.String(), cstr)
	return r
}

func (h *hook) DoMultiCache(c rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) []rueidis.RedisResult {
	before := time.Now()
	for _, m := range multi {
		h.logf("Redis CMD: [%s]", cmdstr(m.Cmd.Commands()))
	}
	rr := c.DoMultiCache(ctx, multi...)
	cost := time.Since(before)
	h.logf("Timing cost for [DoMultiCache]: [%s]", cost.String())
	return rr
}

func (h *hook) Receive(c rueidis.Client, ctx context.Context, sub rueidis.Completed, fn func(msg rueidis.PubSubMessage)) error {
	before := time.Now()
	cstr := cmdstr(sub.Commands())
	err := c.Receive(ctx, sub, fn)
	cost := time.Since(before)
	h.logf("Timing cost [%s] for [Receive]: [%s]", cost.String(), cstr)
	return err
}

func (h *hook) DoStream(c rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	before := time.Now()
	cstr := cmdstr(cmd.Commands())
	res := c.DoStream(ctx, cmd)
	cost := time.Since(before)
	h.logf("Timing cost [%s] for [Do]: [%s]", cost.String(), cstr)
	return res
}

func (h *hook) DoMultiStream(c rueidis.Client, ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	before := time.Now()
	for _, m := range multi {
		h.logf("Redis CMD: [%s]", cmdstr(m.Commands()))
	}
	rr := c.DoMultiStream(ctx, multi...)
	cost := time.Since(before)
	h.logf("Timing cost for [DoMulti]: [%s]", cost.String())
	return rr
}

func cmdstr(cmds []string) string {
	return strings.Join(cmds, " ")
}
