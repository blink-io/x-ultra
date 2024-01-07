package sql

import (
	"context"
	"database/sql"

	"github.com/ilibs/gosql/v2"
)

const (
	AccessorDBW = "dbw"
	RawNameDBW  = "gosql"
)

type (
	idbw = gosql.DB

	DBW struct {
		*idbw
		sqlDB    *sql.DB
		accessor string
		rawName  string
	}
)

var _ HealthChecker = (*DBW)(nil)

func NewDBW(c *Config) (*DBW, error) {
	c = setupConfig(c)
	c.accessor = AccessorDBW

	sqlDB, err := NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	rdb := gosql.OpenWithDB(c.Dialect, sqlDB)

	if c.Logger != nil {
		gosql.SetLogger(PrintfLogger(c.Logger))
	}

	db := &DBW{
		idbw:     rdb,
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDBW,
	}
	return db, nil
}

func (db *DBW) Accessor() string {
	return db.accessor
}

func (db *DBW) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
}
