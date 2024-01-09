package dbp

import (
	"context"
	"database/sql"

	xsql "github.com/blink-io/x/sql"

	"github.com/go-gorp/gorp/v3"
)

const (
	Accessor = "dbp(gorp)"
)

type (
	idb = gorp.DbMap

	DB struct {
		*idb
		sqlDB    *sql.DB
		info     xsql.DBInfo
		accessor string
		rawName  string
	}
)

var _ xsql.HealthChecker = (*DB)(nil)

func New(c *xsql.Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)

	var d gorp.Dialect
	switch c.Dialect {
	case xsql.DialectMySQL:
		d = gorp.MySQLDialect{}
	case xsql.DialectPostgres:
		d = gorp.PostgresDialect{}
	case xsql.DialectSQLite:
		d = gorp.SqliteDialect{}
	default:
		return nil, xsql.ErrUnsupportedDialect
	}

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	opts := applyOptions(ops...)
	if opts != nil {

	}

	rdb := &gorp.DbMap{
		Db:      sqlDB,
		Dialect: d,
	}

	db := &DB{
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
	}
	return db, nil
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
