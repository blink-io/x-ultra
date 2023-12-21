package sql

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/uptrace/bun/schema"
)

type DialectCreator = func(context.Context, ...DOption) schema.Dialect

type DSNCreator = func(*Options) (string, error)

var dialectCreators = make(map[string]DialectCreator)

var dsnCreators = make(map[string]DSNCreator)

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
