package logging

import (
	"context"

	"github.com/blink-io/x/sql/driver/hooks"
)

type Func func(format string, args ...any)

var _ hooks.Hooks = (Func)(nil)

func (f Func) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	f("Executed SQL, query: %s, args: %+v", query, args)
	return ctx, nil
}

func (f Func) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	return ctx, nil
}

type CtxFunc func(ctx context.Context, format string, args ...any)

var _ hooks.Hooks = (CtxFunc)(nil)

func (f CtxFunc) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	f(ctx, "Executed SQL, query: %s, args: %+v", query, args)
	return ctx, nil
}

func (f CtxFunc) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	return ctx, nil
}
