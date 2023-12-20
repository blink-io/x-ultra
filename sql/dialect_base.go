package sql

import (
	"database/sql/driver"
	"time"

	"github.com/uptrace/bun/schema"
)

type DialectFunc = func(...DOption) schema.Dialect

type DSNFunc = func(*Options) (string, error)

var dialectFuncs = make(map[string]DialectFunc)

var dsnFuncs = make(map[string]DSNFunc)

var drivers = make(map[string]driver.Driver)

type dOptions struct {
	loc *time.Location
}

func applyDOptions(ops ...DOption) *dOptions {
	opt := new(dOptions)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

type DOption func(*dOptions)

func WithLocation(loc *time.Location) DOption {
	return func(o *dOptions) {
		o.loc = loc
	}
}
