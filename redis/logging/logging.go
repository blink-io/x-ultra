package logging

import (
	"context"
)

type Logging interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

var _ Logging = (Wrap)(nil)
var _ Logging = (WrapF)(nil)

type Wrap func(context.Context, string, ...interface{})

func (w Wrap) Printf(ctx context.Context, format string, v ...interface{}) {
	w(ctx, format, v...)
}

type WrapF func(string, ...interface{})

func (w WrapF) Printf(ctx context.Context, format string, v ...interface{}) {
	w(format, v...)
}
