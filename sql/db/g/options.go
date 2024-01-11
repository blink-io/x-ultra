package g

import (
	"context"

	"github.com/uptrace/bun"
)

type (
	InsertOption func(*bun.InsertQuery)

	UpdateOption func(*bun.UpdateQuery)

	SelectOption func(*bun.SelectQuery)

	DeleteOption func(*bun.DeleteQuery)

	Where struct {
		q string
		a []any
	}
)

var EmptyWhere = Where{}

type withTxCtxKey struct{}

func WithTx(ctx context.Context) context.Context {
	return context.WithValue(ctx, withTxCtxKey{}, true)
}

func HasTx(ctx context.Context) bool {
	has, ok := ctx.Value(withTxCtxKey{}).(bool)
	return ok && has
}

func InsertIgnore(v bool) InsertOption {
	return func(q *bun.InsertQuery) {
		q.Ignore()
	}
}

func InsertReturning(query string, args ...any) InsertOption {
	return func(q *bun.InsertQuery) {
		q.Returning(query, args...)
	}
}

func UpdateOmitZero(v bool) UpdateOption {
	return func(q *bun.UpdateQuery) {
		if v {
			q.OmitZero()
		}
	}
}

func ForceDelete() DeleteOption {
	return func(q *bun.DeleteQuery) {
		q.ForceDelete()
	}
}

func SelectWhere(es ...Where) SelectOption {
	return func(q *bun.SelectQuery) {
		for _, e := range es {
			q.Where(e.q, e.a...)
		}
	}
}

func SelectColumns(cols ...string) SelectOption {
	return func(q *bun.SelectQuery) {
		q.Column(cols...)
	}
}

func NewWhere(q string, a ...any) Where {
	return Where{q: q, a: a}
}
