package dbx

import (
	"context"
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	"github.com/pocketbase/dbx"
)

const (
	Accessor = "dbx(dbx)"
)

type (
	idb = dbx.DB

	DB struct {
		*idb
		sqlDB    *sql.DB
		accessor string
		rawName  string
		info     xsql.DBInfo
	}
)

var _ xsql.HealthChecker = (*DB)(nil)

func New(c *xsql.Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	opts := applyOptions(ops...)

	rdb := dbx.NewFromDB(sqlDB, c.Dialect)
	rdb.LogFunc = c.Logger
	if opts.queryLogFunc != nil {
		rdb.QueryLogFunc = opts.queryLogFunc
	}
	if opts.queryLogFunc != nil {
		rdb.ExecLogFunc = opts.execLogFunc
	}
	db := &DB{
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
		info:     c.DBInfo(),
	}
	return db, nil
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
