package db

import (
	"time"
)

type (
	dialectOptions struct {
		loc *time.Location
	}

	// DialectOption defines option for dialect
	DialectOption func(*dialectOptions)
)

var (
	dialectors = make(map[string]Dialector)
)

func applyDialectOptions(ops ...DialectOption) *dialectOptions {
	opt := new(dialectOptions)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func DialectWithLoc(loc *time.Location) DialectOption {
	return func(o *dialectOptions) {
		o.loc = loc
	}
}
