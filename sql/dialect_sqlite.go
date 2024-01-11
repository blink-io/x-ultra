package sql

import (
	"context"
	"database/sql/driver"

	//"github.com/glebarez/go-sqlite"
	"github.com/life4/genesis/slices"
	"modernc.org/sqlite"
)

var compatibleSQLiteDialects = []string{
	DialectSQLite,
	"sqlite3",
}

func init() {
	dn := DialectSQLite
	drivers[dn] = GetSQLiteDriver
	dsners[dn] = GetSQLiteDSN
}

func GetSQLiteDSN(dialect string) (Dsner, error) {
	if !IsCompatibleSQLiteDialect(dialect) {
		return nil, ErrUnsupportedDialect
	}
	return func(ctx context.Context, c *Config) (string, error) {
		dsn := c.Host
		return dsn, nil
	}, nil
}

func IsCompatibleSQLiteDialect(dialect string) bool {
	i := slices.FindIndex(compatibleSQLiteDialects, func(i string) bool {
		return i == dialect
	})
	return i > -1
}

func GetSQLiteDriver(dialect string) (driver.Driver, error) {
	if IsCompatibleSQLiteDialect(dialect) {
		return &sqlite.Driver{}, nil
	}
	return nil, ErrUnsupportedDriver
}
