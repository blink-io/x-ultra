package sql

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/go-rel/mysql"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
)

type (
	idbm = rel.Repository
	DBM  struct {
		idbm
	}
)

func NewDBM(o *Options) (*DBM, error) {
	o = setupOptions(o)
	o.accessor = "dbm"

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	var d rel.Adapter
	switch o.Dialect {
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
	if o.Logger != nil {
		rdb.Instrumentation(DBMLogger(o.Logger))
	}
	db := &DBM{
		idbm: rdb,
	}
	return db, nil
}

func DBMLogger(logger func(format string, args ...any)) rel.Instrumenter {
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
