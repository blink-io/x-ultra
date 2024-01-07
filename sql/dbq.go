package sql

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

const (
	AccessorDBQ = "dbq"
	RawNameDBQ  = "goqu"
)

func init() {
	goqu.RegisterDialect(DialectPostgres, postgres.DialectOptions())
	goqu.RegisterDialect(DialectMySQL, mysql.DialectOptions())
	goqu.RegisterDialect(DialectSQLite, sqlite3.DialectOptions())
}

type (
	idbq = goqu.Database

	DBQ struct {
		*idbq
		sqlDB    *sql.DB
		accessor string
		rawName  string
	}
)

var _ HealthChecker = (*DBQ)(nil)

func NewDBQ(c *Config) (*DBQ, error) {
	c = setupConfig(c)
	c.accessor = AccessorDBQ

	sqlDB, err := NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	rdb := goqu.New(c.Dialect, sqlDB)
	if c.Logger != nil {
		rdb.Logger(PrintfLogger(c.Logger))
	}
	if c.Loc != nil {
		goqu.SetTimeLocation(c.Loc)
	}

	db := &DBQ{
		idbq:     rdb,
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDBQ,
	}
	return db, nil
}

func (db *DBQ) Accessor() string {
	return db.accessor
}

func (db *DBQ) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
}
