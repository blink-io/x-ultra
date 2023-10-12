//go:build sqlite

package sql

import (
	"log/slog"

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
	slog.Info("SQLite is enabled")
}

func SQLiteDSN(o *Options) string {
	dsn := o.Host
	return dsn
}
