package dbs

import (
	"github.com/gocraft/dbr/v2"
)

type options struct {
	er dbr.EventReceiver
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}
