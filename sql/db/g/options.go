package g

import (
	"context"
)

type (
	withTxCtxKey struct{}

	insertOptions struct {
		ignore bool

		returning *queryAndArgs
	}

	InsertOption func(*insertOptions)

	updateOptions struct {
		omitZero bool

		returning *queryAndArgs
	}

	UpdateOption func(*updateOptions)

	selectOptions struct {
		cols []string
	}

	SelectOption func(*selectOptions)

	deleteOptions struct {
		force bool
	}

	DeleteOption func(*deleteOptions)

	queryAndArgs struct {
		query string
		args  []any
	}

	Where struct {
		q string
		a []any
	}
)

var EmptyWhere = Where{}

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

func InsertIgnore() InsertOption {
	return func(o *insertOptions) {
		o.ignore = true
	}
}

func InsertReturning(query string, args ...any) InsertOption {
	return func(o *insertOptions) {
		o.returning = &queryAndArgs{
			query: query,
			args:  args,
		}
	}
}

func applyUpdateOptions(ops ...UpdateOption) *updateOptions {
	opts := new(updateOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func UpdateOmitZero() UpdateOption {
	return func(o *updateOptions) {
		o.omitZero = true
	}
}

func UpdateReturning(query string, args ...any) UpdateOption {
	return func(o *updateOptions) {
		o.returning = &queryAndArgs{
			query: query,
			args:  args,
		}
	}
}

func applyDeleteOptions(ops ...DeleteOption) *deleteOptions {
	opts := new(deleteOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DeleteForce() DeleteOption {
	return func(o *deleteOptions) {
		o.force = true
	}
}

func applySelectOptions(ops ...SelectOption) *selectOptions {
	opts := new(selectOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func SelectWhere(ws ...Where) SelectOption {
	return func(o *selectOptions) {
		//for _, e := range es {
		//	q.Where(e.q, e.a...)
		//}
	}
}

func SelectColumns(cols ...string) SelectOption {
	return func(o *selectOptions) {
		o.cols = append(o.cols, cols...)
	}
}
