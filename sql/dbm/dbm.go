package dbm

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
	"time"

	xsql "github.com/blink-io/x/sql"

	"github.com/go-rel/mysql"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
)

const (
	Accessor = "dbm(go-rel)"
)

type (
	idb = rel.Repository
	DB  struct {
		idb
		sqlDB    *sql.DB
		info     xsql.DBInfo
		accessor string
		rawName  string
	}
)

var _ xsql.HealthChecker = (*DB)(nil)

func New(c *xsql.Config) (*DB, error) {
	c = xsql.SetupConfig(c)

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	var d rel.Adapter
	switch c.Dialect {
	case xsql.DialectMySQL:
		d = mysql.New(sqlDB)
	case xsql.DialectPostgres:
		d = postgres.New(sqlDB)
	case xsql.DialectSQLite:
		d = sqlite3.New(sqlDB)
	default:
		return nil, xsql.ErrUnsupportedDialect
	}

	rdb := rel.New(d)
	if c.Logger != nil {
		rdb.Instrumentation(dbmLogger(c.Logger))
	}
	db := &DB{
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
	}
	return db, nil
}

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}

func dbmLogger(logger func(format string, args ...any)) rel.Instrumenter {
	return func(ctx context.Context, op string, message string, args ...interface{}) func(err error) {
		if strings.HasPrefix(op, "rel-") {
			return func(error) {}
		}

		t := time.Now()

		return func(err error) {
			duration := time.Since(t)
			if err != nil {
				msg := fmt.Sprint("[duration: ", duration, " op: ", op, "] ", message, " - ", err)
				logger(msg)
			} else {
				msg := fmt.Sprint("[duration: ", duration, " op: ", op, "] ", message)
				logger(msg)
				slog.Default().Info(msg)
			}
		}
	}
	// no op for rel functions.
}
