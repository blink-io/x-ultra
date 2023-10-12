package buntiming

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type (
	ctxKey struct{}

	QueryHook struct {
	}
)

var _ bun.QueryHook = (*QueryHook)(nil)

func NewQueryHook() *QueryHook {
	h := &QueryHook{}
	return h
}

func (q *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return context.WithValue(ctx, ctxKey{}, time.Now())
}

func (q *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	if before, ok := ctx.Value(ctxKey{}).(time.Time); ok {
		timing := time.Since(before)
		fmt.Println("sql exec timing: ", timing)
	}
}
