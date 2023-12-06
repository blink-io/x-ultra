package sql

import (
	"github.com/blink-io/x/bun/dialect/sqlitedialect"
	"github.com/rqlite/gorqlite/stdlib"
	"github.com/uptrace/bun/schema"
)

const (
	DialectRQLite = "rqlite"
)

func init() {
	dn := DialectRQLite
	drivers[dn] = &stdlib.Driver{}
	dialectFuncs[dn] = func() schema.Dialect {
		return sqlitedialect.New()
	}
	dsnFuncs[dn] = SQLiteDSN
}

func RQLiteDSN(o *Options) string {
	dsn := o.Host
	return dsn
}
