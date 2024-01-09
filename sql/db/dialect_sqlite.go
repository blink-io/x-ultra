//go:build !nosqlite

package db

import (
	"context"

	"github.com/blink-io/x/bun/dialect/sqlitedialect"
	"github.com/blink-io/x/sql"
	xsql "github.com/blink-io/x/sql"

	"github.com/uptrace/bun/schema"
)

func init() {
	dialectors[xsql.DialectSQLite] = NewSQLiteDialect
}

func NewSQLiteDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	dopt := applyDialectOptions(ops...)
	sops := make([]sqlitedialect.Option, 0)
	if dopt.loc != nil {
		sops = append(sops, sqlitedialect.Location(dopt.loc))
	}
	return sqlitedialect.New(sops...)
}

func SQLiteDSN(ctx context.Context, c *sql.Config) (string, error) {
	dsn := c.Host
	return dsn, nil
}
