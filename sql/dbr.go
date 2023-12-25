package sql

import (
	"github.com/blink-io/x/sql/dbr/dialect"

	"github.com/gocraft/dbr/v2"
)

const AccessorDBR = "dbr"

type (
	idbr = dbr.Session

	DBR struct {
		*idbr
		accessor string
	}
)

func NewDBR(o *Options) (*DBR, error) {
	o = setupOptions(o)
	o.accessor = AccessorDBR

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	var d dbr.Dialect
	switch o.Dialect {
	case DialectMySQL:
		d = dialect.MySQL
	case DialectPostgres:
		d = dialect.Postgres
	case DialectSQLite:
		d = dialect.SQLite3
	default:
		return nil, ErrUnsupportedDialect
	}

	cc := &dbr.Connection{
		DB:            sqlDB,
		Dialect:       d,
		EventReceiver: new(dbr.NullEventReceiver),
	}
	rdb := cc.NewSession(nil)
	rdb.Timeout = DefaultTimeout

	db := &DBR{
		idbr:     rdb,
		accessor: o.accessor,
	}
	return db, nil
}

func (d *DBR) Accessor() string {
	return d.accessor
}
