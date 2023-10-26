package logging

import (
	"context"
)

type Logging interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

var _ Logging = (Func)(nil)

type Func func(string, ...interface{})

func (f Func) Printf(ctx context.Context, format string, v ...interface{}) {
	f(format, v...)
}
