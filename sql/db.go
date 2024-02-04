package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"runtime"

	"github.com/blink-io/x/cast"
)

const (
	// DialectMySQL defines MySQL dialect
	DialectMySQL = "mysql"
	// DialectPostgres defines PostgreSQL dialect
	DialectPostgres = "postgres"
	// DialectSQLite defines SQLite dialect. We only support SQLite3 above.
	DialectSQLite = "sqlite"
	// DialectSQLite3 defines SQLite dialect
	DialectSQLite3 = "sqlite3"
)

var (
	ErrUnsupportedDialect = errors.New("unsupported dialect")

	ErrUnsupportedDriver = errors.New("unsupported driver")
)

type (
	Pinger interface {
		PingContext(ctx context.Context) error
	}

	IDB interface {
		Begin() (*sql.Tx, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	}

	WithSqlDB interface {
		SqlDB() *sql.DB
	}

	WithDBInfo interface {
		DBInfo() DBInfo
	}

	HealthChecker interface {
		HealthCheck(context.Context) error
	}

	DBInfo struct {
		Name    string
		Dialect string
	}
)

func NewSqlDB(c *Config) (*sql.DB, error) {
	dialect := GetFormalDialect(c.Dialect)
	ctx := c.Context
	if ctx == nil {
		ctx = context.Background()
	}

	connFn := connectors[dialect]
	if connFn == nil {
		return nil, ErrUnsupportedDialect
	}

	conn, err := connFn(ctx, c)
	if err != nil {
		return nil, err
	}

	var db *sql.DB
	if c.WithOTel {
		otelOps := []OTelOption{
			OTelDBHostPort(hostportToAddr(c.Host, c.Port)),
			OTelDBName(c.Name),
			OTelDBSystem(c.Dialect),
			OTelReportDBStats(),
			OTelAttrs(c.OTelAttrs...),
		}
		if len(c.Accessor) > 0 {
			otelOps = append(otelOps, OTelDBAccessor(c.Accessor))
		}
		db = otelOpenDB(conn, otelOps...)
	} else {
		db = otelWrapper(sql.OpenDB)(conn)
	}

	// Do ping check
	if err := DoPingContext(ctx, db); err != nil {
		return nil, err
	}

	connInitSQL := c.ConnInitSQL
	validationSQL := c.ValidationSQL
	if len(connInitSQL) > 0 {
		if _, err := db.Exec(connInitSQL); err != nil {
			return nil, fmt.Errorf("unable to exec conn_init_sql: %s, reason: %s", connInitSQL, err)
		}
	}
	// Execute validation SQL after bun.DB is initialized
	if len(validationSQL) > 0 {
		if _, err := db.Exec(validationSQL); err != nil {
			return nil, fmt.Errorf("unable to exec validation_sql: %s, reason: %s", validationSQL, err)
		}
	}

	// Reference: https://bun.uptrace.dev/guide/running-bun-in-production.html
	maxIdleConns := c.MaxIdleConns
	maxOpenConns := c.MaxOpenConns
	connMaxLifetime := c.ConnMaxLifetime
	connMaxIdleTime := c.ConnMaxIdleTime
	if maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	} else {
		// TODO In Docker how we should do?
		maxOpenConns = 4 * runtime.GOMAXPROCS(0)
		db.SetMaxOpenConns(maxOpenConns)
	}
	if maxIdleConns > 0 {
		db.SetMaxIdleConns(maxIdleConns)
	} else {
		db.SetMaxIdleConns(maxOpenConns)
	}
	if connMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(connMaxIdleTime)
	}
	if connMaxLifetime > 0 {
		db.SetConnMaxLifetime(connMaxLifetime)
	}

	return db, nil
}

func NewDBInfo(c *Config) DBInfo {
	return DBInfo{
		Name:    c.Name,
		Dialect: c.Dialect,
	}
}

func hostportToAddr(host string, port int) string {
	return net.JoinHostPort(host, cast.ToString(port))
}
