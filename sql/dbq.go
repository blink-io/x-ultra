package sql

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

func init() {
	goqu.RegisterDialect(DialectPostgres, postgres.DialectOptions())
	goqu.RegisterDialect(DialectMySQL, mysql.DialectOptions())
	goqu.RegisterDialect(DialectSQLite, sqlite3.DialectOptions())
}

type idbq = goqu.Database

type DBQ struct {
	*idbq
}

func NewDBQ(o *Options) (*DBQ, error) {
	o = setupOptions(o)
	o.accessor = "dbq"

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	rdb := goqu.New(o.Dialect, sqlDB)
	if o.Logger != nil {
		rdb.Logger(dbqLogger(o.Logger))
	}

	db := &DBQ{
		idbq: rdb,
	}
	return db, nil
}

type dbqLogger func(format string, args ...any)

func (l dbqLogger) Printf(format string, args ...any) {
	l(format, args)
}
