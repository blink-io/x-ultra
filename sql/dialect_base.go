package sql

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/uptrace/bun/schema"
)

type (
	Dsner = func(context.Context, *Options) (string, error)

	Dialector = func(context.Context, ...DOption) schema.Dialect

	dOptions struct {
		loc *time.Location
	}

	DOption func(*dOptions)
)

var (
	drivers = make(map[string]driver.Driver)

	dsnors = make(map[string]Dsner)

	dialectors = make(map[string]Dialector)
)

func applyDOptions(ops ...DOption) *dOptions {
	opt := new(dOptions)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func WithLocation(loc *time.Location) DOption {
	return func(o *dOptions) {
		o.loc = loc
	}
}
