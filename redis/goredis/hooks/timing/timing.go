package timing

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/redis/go-redis/extra/rediscmd/v9"
	"github.com/redis/go-redis/v9"
)

type hook struct {
	logf func(string, ...any)
}

var _ redis.Hook = (*hook)(nil)

func New(ops ...Option) redis.Hook {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	if h.logf == nil {
		h.logf = log.Printf
	}
	return h
}

func (h *hook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		before := time.Now()
		cc, err := next(ctx, network, addr)
		if err == nil {
			cost := time.Since(before)
			h.logf("Timing cost [%s] for DialHook: [%s]", cost.String(), addr)
		}
		return cc, err
	}
}

func (h *hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		before := time.Now()
		err := next(ctx, cmd)
		if err == nil {
			cost := time.Since(before)
			h.logf("Timing cost [%s] for ProcessHook: [%s]", cost.String(), rediscmd.CmdString(cmd))
		}
		return err
	}
}

func (h *hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		before := time.Now()
		err := next(ctx, cmds)
		if err == nil {
			cost := time.Since(before)
			cstr1, cstr2 := rediscmd.CmdsString(cmds)
			h.logf("Timing cost [%s] for ProcessPipelineHook: [%s, %s]", cost.String(), cstr1, cstr2)
		}
		return err
	}
}
