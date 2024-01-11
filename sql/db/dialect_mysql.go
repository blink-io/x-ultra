package db

import (
	"context"

	xsql "github.com/blink-io/x/sql"

	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/schema"
)

func init() {
	dialectors[xsql.DialectMySQL] = NewMySQLDialect
}

func NewMySQLDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	return mysqldialect.New()
}
