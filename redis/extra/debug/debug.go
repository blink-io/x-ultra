package debug

import (
	"context"
	"log"

	"github.com/redis/go-redis/extra/rediscmd/v9"
	"github.com/redis/go-redis/v9"
)

type hook struct {
	log func(string, ...any)
}

var _ redis.Hook = (*hook)(nil)

func New(ops ...Option) redis.Hook {
	h := new(hook)
	for _, o := range ops {
		o(h)
	}
	if h.log == nil {
		h.log = log.Printf
	}
	return h
}

func (h *hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	str := rediscmd.CmdString(cmd)
	h.log("Redis CMD: ", str)
	return ctx, nil
}

func (h *hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	if err := cmd.Err(); err != nil {
	}
	return nil
}

func (h *hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	str1, str2 := rediscmd.CmdsString(cmds)
	h.log("%s: %s", str1, str2)
	return ctx, nil
}

func (h *hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
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
