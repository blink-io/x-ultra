package bunlogging

import (
	"context"

	"github.com/uptrace/bun"
)

var _ bun.QueryHook = (*QueryHook)(nil)

type QueryHook struct {
}

func (q *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {

	return ctx
}

func (q *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
}
