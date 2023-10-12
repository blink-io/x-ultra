//go:build sqlite3

package sql

import (
	"database/sql/driver"

	"github.com/blink-io/x/bun/dialect/sqlitedialect"

	"github.com/uptrace/bun/schema"
	"modernc.org/sqlite"
	sqlite3lib "modernc.org/sqlite/lib"
)

const (
	DialectSQLite = "sqlite3"
)

func init() {
	fn := func() schema.Dialect {
		return sqlitedialect.New()
	}
	SetDialectFn(DialectSQLite, fn)
	SetDriverFn(DialectSQLite, GetSqlite3Driver)
}

func Sqlite3DSN(o *Options) string {
	dsn := o.Host
	return dsn
}

func GetSqlite3Driver(o *Options) (string, driver.Driver) {
	dsn := Sqlite3DSN(o)
	drv := &sqlite.Driver{}
	return dsn, drv
}

func IsSQLiteConstraintCodes(code int) bool {
	return sqlite3lib.SQLITE_CONSTRAINT == code ||
		sqlite3lib.SQLITE_CONSTRAINT_PRIMARYKEY == code ||
		sqlite3lib.SQLITE_CONSTRAINT_UNIQUE == code ||
		sqlite3lib.SQLITE_CONSTRAINT_ROWID == code ||
		sqlite3lib.SQLITE_CONSTRAINT_FOREIGNKEY == code ||
		sqlite3lib.SQLITE_CONSTRAINT_NOTNULL == code
}
