package logging

import (
	"context"

	"github.com/blink-io/x/sql/hooks"
)

type Func func(string, ...any)

var _ hooks.Hooks = (Func)(nil)

func (f Func) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	f(query, args...)
	return ctx, nil
}

func (f Func) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	return ctx, nil
}

type CtxFunc func(context.Context, string, ...any)

var _ hooks.Hooks = (CtxFunc)(nil)

func (f CtxFunc) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	f(ctx, query, args...)
	return ctx, nil
}

func (f CtxFunc) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	return ctx, nil
}
