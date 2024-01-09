package dbq

import (
	"context"
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

const (
	Accessor = "dbq(goqu)"
)

func init() {
	goqu.RegisterDialect(xsql.DialectPostgres, postgres.DialectOptions())
	goqu.RegisterDialect(xsql.DialectMySQL, mysql.DialectOptions())
	goqu.RegisterDialect(xsql.DialectSQLite, sqlite3.DialectOptions())
}

type (
	idb = goqu.Database

	Config = xsql.Config

	DB struct {
		*idb
		sqlDB    *sql.DB
		accessor string
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

	if c.Loc != nil {
		goqu.SetTimeLocation(c.Loc)
	}

	opts := applyOptions(ops...)
	if opts != nil {
	}

	rdb := goqu.New(c.Dialect, sqlDB)
	if c.Logger != nil {
		rdb.Logger(xsql.PrintfLogger(c.Logger))
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
