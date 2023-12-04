package logging

import (
	"context"
)

// Logging is copied from github.com/redis/go-redis/internal/log
type Logging interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

var _ Logging = (Func)(nil)

type Func func(string, ...any)

func (f Func) Printf(ctx context.Context, format string, v ...any) {
	f(format, v...)
}

type CtxFunc func(context.Context, string, ...any)

func (f CtxFunc) Printf(ctx context.Context, format string, v ...any) {
	f(ctx, format, v...)
}
