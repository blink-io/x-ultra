package g

import (
	"context"

	"github.com/uptrace/bun/schema"
)

type (
	queryWithArgs = schema.QueryWithArgs

	withTxCtxKey struct{}

	insertOptions struct {
		ignore bool

		returning *schema.QueryWithArgs
	}

	InsertOption func(*insertOptions)

	updateOptions struct {
		omitZero bool

		returning *queryWithArgs
	}

	UpdateOption func(*updateOptions)

	selectOptions struct {
		cols    []string
		where   []*queryWithArgs
		whereOr []*queryWithArgs
	}

	SelectOption func(*selectOptions)

	deleteOptions struct {
		force bool

		returning *queryWithArgs
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

func InsertIgnore() InsertOption {
	return func(o *insertOptions) {
		o.ignore = true
	}
}

func InsertReturning(query string, args ...any) InsertOption {
	return func(o *insertOptions) {
		o.returning = &queryWithArgs{
			Query: query,
			Args:  args,
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
		o.returning = &queryWithArgs{
			Query: query,
			Args:  args,
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

func DeleteReturning(query string, args ...any) DeleteOption {
	return func(o *deleteOptions) {
		o.returning = &queryWithArgs{
			Query: query,
			Args:  args,
		}
	}
}

func applySelectOptions(ops ...SelectOption) *selectOptions {
	opts := new(selectOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func SelectWhere(query string, args ...any) SelectOption {
	return func(o *selectOptions) {
		o.where = append(
			o.where,
			&queryWithArgs{
				Query: query,
				Args:  args,
			},
		)
	}
}

func SelectWhereOr(query string, args ...any) SelectOption {
	return func(o *selectOptions) {
		o.whereOr = append(
			o.whereOr,
			&queryWithArgs{
				Query: query,
				Args:  args,
			},
		)
	}
}

func SelectColumns(cols ...string) SelectOption {
	return func(o *selectOptions) {
		o.cols = append(o.cols, cols...)
	}
}
