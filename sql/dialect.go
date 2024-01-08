package sql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
	"runtime"

	"github.com/blink-io/x/cast"
	"github.com/blink-io/x/sql/driver/hooks"

	"github.com/uptrace/bun/schema"
)

var (
	ErrUnsupportedDialect = errors.New("unsupported dialect")

	ErrUnsupportedDriver = errors.New("unsupported driver")
)

func GetDialect(c *Config, ops ...DialectOption) (schema.Dialect, *sql.DB, error) {
	c = setupConfig(c)

	ctx := c.Context
	dialect := c.Dialect

	var sd schema.Dialect
	if dfn, ok := dialectors[dialect]; ok {
		sd = dfn(ctx, ops...)
	} else {
		return nil, nil, ErrUnsupportedDialect
	}

	db, err := NewSqlDB(c)
	if err != nil {
		return nil, nil, err
	}

	return sd, db, nil
}

func NewSqlDB(c *Config) (*sql.DB, error) {
	dialect := c.Dialect
	ctx := c.Context

	var dsn string
	var err error
	if dfn, ok := dsners[dialect]; ok {
		dsn, err = dfn(ctx, c)
		c.dsn = dsn
		if err != nil {
			return nil, err
		}
	} else {
		return nil, ErrUnsupportedDialect
	}

	var drv driver.Driver
	if dd, ok := drivers[dialect]; ok {
		drv = dd
	} else {
		return nil, ErrUnsupportedDriver
	}

	driverHooks := c.DriverHooks
	if len(driverHooks) > 0 {
		drv = hooks.Wrap(drv, hooks.Compose(driverHooks...))
	}

	conn := &dsnConnector{dsn: dsn, driver: drv}
	hostPort := net.JoinHostPort(c.Host, cast.ToString(c.Port))
	var db *sql.DB
	if c.WithOTel {
		otelOps := []OTelOption{
			OTelDBName(c.Name),
			OTelDBSystem(c.Dialect),
			OTelDBHostPort(hostPort),
			OTelReportDBStats(),
			OTelAttrs(c.Attrs...),
		}
		if len(c.accessor) > 0 {
			otelOps = append(otelOps, OTelDBAccessor(c.accessor))
		}
		db = otelOpenDB(conn, otelOps...)
	} else {
		db = sqlOpenDB(conn)
	}

	// Ignore driver.ErrSkip when the Conn does not implement driver.Pinger interface
	if err := db.Ping(); err != nil && !errors.Is(err, driver.ErrSkip) {
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
