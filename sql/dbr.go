package sql

import (
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
)

type (
	dbrs = dbr.Session

	DBR struct {
		*dbrs
	}
)

func NewDBR(o *Options) (*DBR, error) {
	o = setupOptions(o)
	o.accessor = "dbr"

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	var d dbr.Dialect
	switch o.Dialect {
	case DialectMySQL:
		d = dialect.MySQL
	case DialectPostgres:
		d = dialect.PostgreSQL
	case DialectSQLite:
		d = dialect.SQLite3
	case DialectMSSQL:
		d = dialect.MSSQL
	default:
		return nil, ErrUnsupportedDialect
	}

	cc := &dbr.Connection{
		DB:            sqlDB,
		Dialect:       d,
		EventReceiver: new(dbr.NullEventReceiver),
	}
	ss := cc.NewSession(nil)
	ss.Timeout = DefaultTimeout

	db := &DBR{
		dbrs: ss,
	}
	return db, nil
}
