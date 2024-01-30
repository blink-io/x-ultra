package dbx

import (
	"context"
	"database/sql"
	"io"

	xsql "github.com/blink-io/x/sql"
	"github.com/pocketbase/dbx"
)

const (
	Accessor = "dbx(dbx)"
)

type (
	idb = dbx.DB

	IDB interface {
		DBF

		io.Closer

		xsql.IDBExt
	}

	DB struct {
		*idb
		sqlDB    *sql.DB
		accessor string
		rawName  string
		info     xsql.DBInfo
	}
)

var _ IDB = (*DB)(nil)

func New(c *xsql.Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)
	c.Accessor = Accessor

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	opts := applyOptions(ops...)
	// Setup dbTag
	if dbTag := opts.dbTag; len(dbTag) > 0 {
		dbx.DbTag = dbTag
	}

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

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}

func (db *DB) DBInfo() xsql.DBInfo {
	return db.info
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
