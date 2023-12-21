package sql

import (
	"context"
	"errors"

	"github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/schema"
)

const (
	DialectMSSQL = "mssql"
)

func init() {
	dn := DialectMSSQL
	drivers[dn] = &mssql.Driver{}
	dialectors[dn] = func(ctx context.Context, ops ...DOption) schema.Dialect {
		return mssqldialect.New()
	}
	dsnors[dn] = MSSQLDSN
}

func MSSQLDSN(ctx context.Context, o *Options) (string, error) {
	return "", errors.New("unsupported for MSSQL")
}
