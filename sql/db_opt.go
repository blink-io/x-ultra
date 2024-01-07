package sql

import (
	"time"
)

type dbOptions struct {
	queryHooks []QueryHook
	loc        *time.Location
}

type DBOption func(*dbOptions)

func applyDBOptions(ops ...DBOption) *dbOptions {
	opt := new(dbOptions)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func DBLoc(loc *time.Location) DBOption {
	return func(o *dbOptions) {
		o.loc = loc
	}
}

func DBQueryHook(h QueryHook) DBOption {
	return func(o *dbOptions) {
		o.queryHooks = append(o.queryHooks, h)
	}
}
