package requestid

import "context"

type contextKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, contextKey{}, requestID)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (requestID string) {
	requestID, _ = ctx.Value(contextKey{}).(string)
	return
}
