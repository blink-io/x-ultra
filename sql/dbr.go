package sql

import (
	"context"
	"database/sql"

	"github.com/blink-io/x/sql/dbr/dialect"

	"github.com/gocraft/dbr/v2"
)

const (
	AccessorDBR = "dbr"
	RawNameDBR  = "dbr"
)

type (
	idbr = dbr.Session

	DBR struct {
		*idbr
		sqlDB    *sql.DB
		accessor string
		rawName  string
		info     DBInfo
	}
)

var _ HealthChecker = (*DBR)(nil)

func NewDBR(c *Config) (*DBR, error) {
	c = setupConfig(c)
	c.accessor = AccessorDBR

	sqlDB, err := NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	var d dbr.Dialect
	switch c.Dialect {
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
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDBR,
		info:     newDBInfo(c),
	}
	return db, nil
}

func (db *DBR) Accessor() string {
	return db.accessor
}

func (db *DBR) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
}
