package bun

import (
	"github.com/uptrace/bun"
)

type QueryHook = bun.QueryHook

type options struct {
	queryHooks []QueryHook
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithQueryHooks(hooks ...QueryHook) Option {
	return func(o *options) {
		o.queryHooks = append(o.queryHooks, hooks...)
	}
}
