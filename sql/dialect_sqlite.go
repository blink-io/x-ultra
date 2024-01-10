//go:build !nosqlite

package sql

import (
	"context"
	"database/sql/driver"

	"github.com/glebarez/go-sqlite"
	"github.com/life4/genesis/slices"
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
	i, _ := slices.Index(compatibleSQLiteDialects, dialect)
	return i > 0
}

func GetSQLiteDriver(dialect string) driver.Driver {
	if IsCompatibleSQLiteDialect(dialect) {
		return &sqlite.Driver{}
	}
	return nil
}
