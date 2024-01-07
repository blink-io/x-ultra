package sql

import (
	"context"
	"database/sql"

	"github.com/go-gorp/gorp/v3"
)

const (
	AccessorDBP = "dbp"
	RawNameDBP  = "gorp"
)

type (
	idbp = gorp.DbMap

	DBP struct {
		*idbp
		sqlDB    *sql.DB
		info     DBInfo
		accessor string
		rawName  string
	}
)

var _ HealthChecker = (*DBP)(nil)

func NewDBP(c *Config) (*DBP, error) {
	c = setupConfig(c)
	c.accessor = AccessorDBP

	var d gorp.Dialect
	switch c.Dialect {
	case DialectMySQL:
		d = gorp.MySQLDialect{}
	case DialectPostgres:
		d = gorp.PostgresDialect{}
	case DialectSQLite:
		d = gorp.SqliteDialect{}
	default:
		return nil, ErrUnsupportedDialect
	}

	sqlDB, err := NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	rdb := &gorp.DbMap{
		Db:      sqlDB,
		Dialect: d,
	}

	db := &DBP{
		idbp:     rdb,
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDBP,
	}
	return db, nil
}

func (db *DBP) Accessor() string {
	return db.accessor
}

func (db *DBP) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
}
