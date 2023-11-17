package sql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
	"runtime"

	"github.com/blink-io/x/cast"
	"github.com/blink-io/x/sql/hooks"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

func GetDialect(o *Options) (schema.Dialect, *sql.DB, error) {
	o, err := setupOptions(o)
	if err != nil {
		return nil, nil, err
	}

	dialect := o.Dialect

	var dsn string
	if dsnFunc, ok := dsnFuncs[dialect]; ok {
		dsn = dsnFunc(o)
	} else {
		return nil, nil, fmt.Errorf("unsupoorted dsn for dialect: %s", dialect)
	}

	var sd schema.Dialect
	if dlaFunc, ok := dialectFuncs[dialect]; ok {
		sd = dlaFunc()
	} else {
		return nil, nil, fmt.Errorf("unsupoorted dialect: %s", dialect)
	}

	var drv driver.Driver
	if dd, ok := drivers[dialect]; ok {
		drv = dd
	} else {
		return nil, nil, fmt.Errorf("unsupoorted driver for dialect: %s", dialect)
	}

	driverHooks := o.DriverHooks
	if len(driverHooks) > 0 {
		drv = hooks.Wrap(drv, hooks.Compose(driverHooks...))
	}

	conn := &dsnConnector{dsn: dsn, driver: drv}
	hostPort := net.JoinHostPort(o.Host, cast.ToString(o.Port))
	var db *sql.DB
	if o.WithOTel {
		db = otelOpenDB(conn,
			OTelDBName(o.Name),
			OTelDBSystem(o.Dialect),
			OTelDBAccessMethod("bun"+" "+bun.Version()),
			OTelDBHostPort(hostPort),
			OTelReportDBStats(),
		)
	} else {
		db = sqlOpenDB(conn)
	}

	// Ignore driver.ErrSkip when the Conn does not implement driver.Pinger interface
	if err := db.Ping(); err != nil && !errors.Is(err, driver.ErrSkip) {
		return nil, nil, err
	}

	connInitSQL := o.ConnInitSQL
	validationSQL := o.ValidationSQL
	if len(connInitSQL) > 0 {
		if _, err := db.Exec(connInitSQL); err != nil {
			return nil, nil, fmt.Errorf("unable to exec conn_init_sql: %s, reason: %s", connInitSQL, err)
		}
	}
	// Execute validation SQL after bun.DB is initialized
	if len(validationSQL) > 0 {
		if _, err := db.Exec(validationSQL); err != nil {
			return nil, nil, fmt.Errorf("unable to exec validation_sql: %s, reason: %s", validationSQL, err)
		}
	}

	// Reference: https://bun.uptrace.dev/guide/running-bun-in-production.html
	maxIdleConns := o.MaxIdleConns
	maxOpenConns := o.MaxOpenConns
	connMaxLifetime := o.ConnMaxLifetime
	connMaxIdleTime := o.ConnMaxIdleTime
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

	return sd, db, err
}
