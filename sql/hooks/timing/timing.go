package timing

import (
	"context"
	"log/slog"
	"time"

	"github.com/blink-io/x/sql/hooks"
)

type ctxKey struct{}

type hook struct {
}

func New() hooks.Hooks {

	return &hook{}
}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	newCtx := context.WithValue(ctx, ctxKey{}, time.Now())
	return newCtx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	if before, ok := ctx.Value(ctxKey{}).(time.Time); ok {
		d := time.Since(before)
		slog.Info("timing", slog.Duration("cost", d))
	}
	return ctx, nil
}
