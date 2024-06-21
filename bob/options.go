package bob

import (
	"github.com/life4/genesis/slices"
)

type options struct {
	wrappers []ExecWrapper
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func ExecWrappers(ws ...ExecWrapper) Option {
	return func(o *options) {
		o.wrappers = slices.Concat(o.wrappers, ws)
	}
}
