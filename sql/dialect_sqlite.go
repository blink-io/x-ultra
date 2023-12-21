//go:build !nosqlite

package sql

import (
	"context"

	"github.com/blink-io/x/bun/dialect/sqlitedialect"

	"github.com/glebarez/go-sqlite"
	"github.com/uptrace/bun/schema"
)

const (
	DialectSQLite = "sqlite"
)

func init() {
	dn := DialectSQLite
	drivers[dn] = &sqlite.Driver{}
	dialectCreators[dn] = func(ctx context.Context, ops ...DOption) schema.Dialect {
		dopt := applyDOptions(ops...)
		sops := make([]sqlitedialect.Option, 0)
		if dopt.loc != nil {
			sops = append(sops, sqlitedialect.Location(dopt.loc))
		}
		return sqlitedialect.New(sops...)
	}
	dsnCreators[dn] = SQLiteDSN
}

func SQLiteDSN(o *Options) (string, error) {
	dsn := o.Host
	return dsn, nil
}
