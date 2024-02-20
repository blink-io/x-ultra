package dbz

import (
	"context"
	"database/sql"
	"io"

	xsql "github.com/blink-io/x/sql"
)

const (
	Accessor = "dbz(todo)"
)

type (
	idb = *sql.DB

	IDB interface {
		DBF

		io.Closer

		xsql.IDBExt
	}

	DB struct {
		idb
		sqlDB *sql.DB
		info  xsql.DBInfo
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
	if opts != nil {

	}

	s := &DB{
		//idb:   rdb,
		sqlDB: sqlDB,
		info:  xsql.NewDBInfo(c),
	}

	return s, nil
}

func (db *DB) Close() error {
	return nil
}

func (db *DB) DBInfo() xsql.DBInfo {
	return db.info
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}
