package bun

import (
	"context"
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	"github.com/uptrace/bun/schema"
)

type (
	Dialector = func(context.Context, ...DialectOption) schema.Dialect
)

func GetDialect(c *xsql.Config, ops ...DialectOption) (schema.Dialect, *sql.DB, error) {
	c = xsql.SetupConfig(c)

	ctx := c.Context
	dialect := xsql.GetFormalDialect(c.Dialect)

	var sd schema.Dialect
	if dfn, ok := dialectors[dialect]; ok {
		sd = dfn(ctx, ops...)
	} else {
		return nil, nil, xsql.ErrUnsupportedDialect
	}

	db, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, nil, err
	}

	return sd, db, nil
}
