package sql

import (
	"context"
	"database/sql"

	"github.com/pocketbase/dbx"
)

const (
	AccessorDBX = "dbx"
	RawNameDBX  = "dbx"
)

type (
	idbx = dbx.DB

	DBX struct {
		*idbx
		sqlDB    *sql.DB
		accessor string
		rawName  string
		info     DBInfo
	}
)

var _ HealthChecker = (*DBX)(nil)

func NewDBX(c *Config) (*DBX, error) {
	c = setupConfig(c)
	c.accessor = AccessorDBX

	sqlDB, err := NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	rdb := dbx.NewFromDB(sqlDB, c.Dialect)
	rdb.LogFunc = c.Logger
	db := &DBX{
		idbx:     rdb,
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDBX,
		info:     newDBInfo(c),
	}
	return db, nil
}

func (db *DBX) Accessor() string {
	return db.accessor
}

func (db *DBX) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
}
