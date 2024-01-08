package sql

import (
	"github.com/gocraft/dbr/v2"
)

type dbrOptions struct {
	er dbr.EventReceiver
}

type DBROption func(*dbrOptions)

func applyDBROptions(ops ...DBROption) *dbrOptions {
	opts := new(dbrOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DBREventReceiver(er dbr.EventReceiver) DBROption {
	return func(o *dbrOptions) {
		o.er = er
	}
}
