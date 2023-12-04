//go:build !nosqlite

package sql

import (
	"github.com/blink-io/x/bun/dialect/sqlitedialect"

	"github.com/uptrace/bun/schema"
	"modernc.org/sqlite"
)

const (
	DialectSQLite = "sqlite"
)

func init() {
	dn := DialectSQLite
	drivers[dn] = &sqlite.Driver{}
	dialectFuncs[dn] = func() schema.Dialect {
		return sqlitedialect.New()
	}
	dsnFuncs[dn] = SQLiteDSN
}

func SQLiteDSN(o *Options) string {
	dsn := o.Host
	return dsn
}
