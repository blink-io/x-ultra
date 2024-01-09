package dbw

import (
	"context"
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	"github.com/ilibs/gosql/v2"
)

const (
	Accessor = "dbw(gosql)"
)

type (
	idb = gosql.DB

	DB struct {
		*idb
		sqlDB    *sql.DB
		accessor string
		info     xsql.DBInfo
	}
)

var _ xsql.HealthChecker = (*DB)(nil)

func New(c *xsql.Config) (*DB, error) {
	c = xsql.SetupConfig(c)

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	rdb := gosql.OpenWithDB(c.Dialect, sqlDB)

	if c.Logger != nil {
		gosql.SetLogger(xsql.PrintfLogger(c.Logger))
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
