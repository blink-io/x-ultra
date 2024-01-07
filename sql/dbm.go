package sql

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/go-rel/mysql"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
)

const (
	AccessorDBM = "dbm"
	RawNameDBM  = "go-rel"
)

type (
	idbm = rel.Repository
	DBM  struct {
		idbm
		sqlDB    *sql.DB
		info     DBInfo
		accessor string
		rawName  string
	}
)

var _ HealthChecker = (*DBM)(nil)

func NewDBM(c *Config) (*DBM, error) {
	c = setupConfig(c)
	c.accessor = AccessorDBM

	sqlDB, err := NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	var d rel.Adapter
	switch c.Dialect {
	case DialectMySQL:
		d = mysql.New(sqlDB)
	case DialectPostgres:
		d = postgres.New(sqlDB)
	case DialectSQLite:
		d = sqlite3.New(sqlDB)
	default:
		return nil, ErrUnsupportedDialect
	}

	rdb := rel.New(d)
	if c.Logger != nil {
		rdb.Instrumentation(dbmLogger(c.Logger))
	}
	db := &DBM{
		idbm:     rdb,
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDBM,
	}
	return db, nil
}

func (db *DBM) Accessor() string {
	return db.accessor
}

func (db *DBM) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
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
