package dbs

import (
	"context"
	"database/sql"

	xsql "github.com/blink-io/x/sql"
)

const (
	Accessor = "dbs(squirrel)"
)

type (
	IDB interface {
		xsql.IDBExt
	}

	DB struct {
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

	s := &DB{
		sqlDB: sqlDB,
		info:  xsql.NewDBInfo(c),
	}

	return s, nil
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
