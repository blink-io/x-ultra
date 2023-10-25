package realip

import (
	"context"
)

type contextKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, contextKey{}, ip)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (ip string) {
	ip, _ = ctx.Value(contextKey{}).(string)
	return
}
