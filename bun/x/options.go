package x

import (
	"context"

	rdb "github.com/blink-io/x/bun"
)

type (
	withTxCtxKey struct{}

	insertOptions struct {
		ignore         bool
		columns        []string
		excludeColumns []string
		columnExprs    []*rdb.QueryWithArgs
		returning      *rdb.QueryWithArgs
	}

	InsertOption func(*insertOptions)

	updateOptions struct {
		omitZero       bool
		bulk           bool
		columns        []string
		excludeColumns []string
		FQN            string
		forceIndexes   []string
		ignoreIndexes  []string
		modelTableExpr *rdb.QueryWithArgs
		returning      *rdb.QueryWithArgs
	}

	UpdateOption func(*updateOptions)

	selectOptions struct {
		queryFunc      func(*rdb.SelectQuery) *rdb.SelectQuery
		queryBuilder   func(rdb.QueryBuilder) rdb.QueryBuilder
		distinct       bool
		distinctOn     []*rdb.QueryWithArgs
		limit          int
		offset         int
		columns        []string
		excludeColumns []string
		columnExprs    []*rdb.QueryWithArgs
		orders         []string
		where          []*rdb.QueryWithArgs
		whereOr        []*rdb.QueryWithArgs
	}

	SelectOption func(*selectOptions)

	deleteOptions struct {
		forceDelete bool
		returning   *rdb.QueryWithArgs
		where       []*rdb.QueryWithArgs
		whereOr     []*rdb.QueryWithArgs
	}

	DeleteOption func(*deleteOptions)
)

func WithTx(ctx context.Context) context.Context {
	return context.WithValue(ctx, withTxCtxKey{}, true)
}

func HasTx(ctx context.Context) bool {
	has, ok := ctx.Value(withTxCtxKey{}).(bool)
	return ok && has
}

func applyInsertOptions(ops ...InsertOption) *insertOptions {
	opts := new(insertOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithInsertIgnore() InsertOption {
	return func(o *insertOptions) {
		o.ignore = true
	}
}

func WithInsertReturning(query string, args ...any) InsertOption {
	return func(o *insertOptions) {
		o.returning = safeQuery(query, args...)
	}
}

func applyUpdateOptions(ops ...UpdateOption) *updateOptions {
	opts := new(updateOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithUpdateOmitZero() UpdateOption {
	return func(o *updateOptions) {
		o.omitZero = true
	}
}

func WithUpdateReturning(query string, args ...any) UpdateOption {
	return func(o *updateOptions) {
		o.returning = safeQuery(query, args...)
	}
}

func applyDeleteOptions(ops ...DeleteOption) *deleteOptions {
	opts := new(deleteOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithDeleteForce() DeleteOption {
	return func(o *deleteOptions) {
		o.forceDelete = true
	}
}

func WithDeleteReturning(query string, args ...any) DeleteOption {
	return func(o *deleteOptions) {
		o.returning = safeQuery(query, args...)
	}
}

func WithDeleteWhere(query string, args ...any) DeleteOption {
	return func(o *deleteOptions) {
		o.where = append(
			o.where,
			safeQuery(query, args...),
		)
	}
}

func WithDeleteWhereOr(query string, args ...any) DeleteOption {
	return func(o *deleteOptions) {
		o.whereOr = append(
			o.whereOr,
			safeQuery(query, args...),
		)
	}
}

func applySelectOptions(ops ...SelectOption) *selectOptions {
	opts := new(selectOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithSelectLimit(limit int) SelectOption {
	return func(o *selectOptions) {
		o.limit = limit
	}
}

func WithSelectOffset(offset int) SelectOption {
	return func(o *selectOptions) {
		o.offset = offset
	}
}

func WithSelectWhere(query string, args ...any) SelectOption {
	return func(o *selectOptions) {
		o.where = append(
			o.where,
			safeQuery(query, args...),
		)
	}
}

func WithSelectWhereOr(query string, args ...any) SelectOption {
	return func(o *selectOptions) {
		o.whereOr = append(
			o.whereOr,
			safeQuery(query, args...),
		)
	}
}

func WithSelectColumns(columns ...string) SelectOption {
	return func(o *selectOptions) {
		o.columns = append(o.columns, columns...)
	}
}

func WithSelectOrders(orders ...string) SelectOption {
	return func(o *selectOptions) {
		o.orders = append(o.orders, orders...)
	}
}

func WithSelectQuery(queryFunc func(*rdb.SelectQuery) *rdb.SelectQuery) SelectOption {
	return func(o *selectOptions) {
		o.queryFunc = queryFunc
	}
}

func safeQuery(query string, args ...any) *rdb.QueryWithArgs {
	sq := rdb.SafeQuery(query, args)
	return &sq
}
