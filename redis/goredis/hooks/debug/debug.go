package debug

import (
	"context"
	"log"
	"net"

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
		return next(ctx, network, addr)
	}
}

func (h *hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		str := rediscmd.CmdString(cmd)
		h.logf("Redis CMD: ", str)
		return next(ctx, cmd)
	}
}

func (h *hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		str1, str2 := rediscmd.CmdsString(cmds)
		h.logf("%s: %s", str1, str2)
		return next(ctx, cmds)
	}
}
