package logging

import (
	"context"
)

type Logging interface {
	Printf(ctx context.Context, format string, v ...any)
}

var _ Logging = (Func)(nil)

type Func func(string, ...any)

func (f Func) Printf(ctx context.Context, format string, v ...any) {
	f(format, v...)
}
