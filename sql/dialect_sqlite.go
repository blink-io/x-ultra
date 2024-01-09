//go:build !nosqlite

package sql

import (
	"context"

	"github.com/glebarez/go-sqlite"
)

func init() {
	dn := DialectSQLite
	drivers[dn] = &sqlite.Driver{}
	dsners[dn] = SQLiteDSN
}

func SQLiteDSN(ctx context.Context, c *Config) (string, error) {
	dsn := c.Host
	return dsn, nil
}
