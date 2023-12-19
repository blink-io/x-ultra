package sql

import (
	"database/sql/driver"

	"github.com/uptrace/bun/schema"
)

type DialectFunc = func() schema.Dialect

type DSNFunc = func(*Options) (string, error)

var dialectFuncs = make(map[string]DialectFunc)

var dsnFuncs = make(map[string]DSNFunc)

var drivers = make(map[string]driver.Driver)
