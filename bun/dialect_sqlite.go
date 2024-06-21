//go:build !nosqlite

package bun

import (
	"context"

	"github.com/blink-io/x/bun/dialect/sqlitedialect"
	xsql "github.com/blink-io/x/sql"

	"github.com/uptrace/bun/schema"
)

func init() {
	dialectors[xsql.DialectSQLite] = NewSQLiteDialect
}

func NewSQLiteDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	dopts := applyDialectOptions(ops...)
	sops := make([]sqlitedialect.Option, 0)
	if dopts.loc != nil {
		sops = append(sops, sqlitedialect.Location(dopts.loc))
	}
	return sqlitedialect.New(sops...)
}
