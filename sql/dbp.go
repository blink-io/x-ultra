package sql

import (
	"github.com/go-gorp/gorp/v3"
)

const (
	AccessorDBP = "dbp"
)

type (
	idbp = gorp.DbMap

	DBP struct {
		*idbp
		accessor string
	}
)

func NewDBP(o *Options) (*DBP, error) {
	o = setupOptions(o)
	o.accessor = AccessorDBP

	var d gorp.Dialect
	switch o.Dialect {
	case DialectMySQL:
		d = gorp.MySQLDialect{}
	case DialectPostgres:
		d = gorp.PostgresDialect{}
	case DialectSQLite:
		d = gorp.SqliteDialect{}
	default:
		return nil, ErrUnsupportedDialect
	}

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	rdb := &gorp.DbMap{
		Db:      sqlDB,
		Dialect: d,
	}

	db := &DBP{
		idbp:     rdb,
		accessor: o.accessor,
	}
	return db, nil
}

func (d *DBP) Accessor() string {
	return d.accessor
}
