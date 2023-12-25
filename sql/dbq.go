package sql

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

const (
	AccessorDBQ = "dbq"
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
		accessor string
	}
)

func NewDBQ(o *Options) (*DBQ, error) {
	o = setupOptions(o)
	o.accessor = AccessorDBQ

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	rdb := goqu.New(o.Dialect, sqlDB)
	if o.Logger != nil {
		rdb.Logger(PrintfLogger(o.Logger))
	}
	if o.Loc != nil {
		goqu.SetTimeLocation(o.Loc)
	}

	db := &DBQ{
		idbq:     rdb,
		accessor: o.accessor,
	}
	return db, nil
}

func (d *DBQ) Accessor() string {
	return d.accessor
}
