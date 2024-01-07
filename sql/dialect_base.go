package sql

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/uptrace/bun/schema"
)

type (
	Dsner = func(context.Context, *Config) (string, error)

	Dialector = func(context.Context, ...DialectOption) schema.Dialect

	dialectOptions struct {
		loc *time.Location
	}

	// DialectOption defines option for dialect
	DialectOption func(*dialectOptions)
)

var (
	drivers = make(map[string]driver.Driver)

	dsners = make(map[string]Dsner)

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
