package session

import (
	"context"
)

type sessMgrCtxKey struct{}

func NewContext(ctx context.Context, m Manager) context.Context {
	return context.WithValue(ctx, sessMgrCtxKey{}, m)
}

func FromContext(ctx context.Context) (Manager, bool) {
	m, ok := ctx.Value(sessMgrCtxKey{}).(Manager)
	return m, ok
}
