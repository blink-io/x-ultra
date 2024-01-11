package db

import (
	"context"

	xsql "github.com/blink-io/x/sql"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/schema"
)

func init() {
	dialectors[xsql.DialectPostgres] = NewPostgresDialect
}

func NewPostgresDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	return pgdialect.New()
}
