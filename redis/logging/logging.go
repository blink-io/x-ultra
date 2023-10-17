package logging

import (
	"context"
)

type Logging interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

var _ Logging = (Fn)(nil)

type Fn func(string, ...interface{})

func (f Fn) Printf(ctx context.Context, format string, v ...interface{}) {
	f(format, v...)
}
