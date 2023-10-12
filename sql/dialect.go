package sql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"runtime"

	"github.com/blink-io/x/sql/hooks"

	"github.com/uptrace/bun/schema"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/attribute"
)

func GetDialect(o *Options) (schema.Dialect, *sql.DB, error) {
	o, err := setupOptions(o)
	if err != nil {
		return nil, nil, err
	}

	dialect := o.Dialect
	name := o.Name

	dltFn, err := GetDialectFn(dialect)
	if err != nil {
		return nil, nil, err
	}
	dlt := dltFn()

	drvFn, err := GetDriverFn(dialect)
	if err != nil {
		return nil, nil, err
	}
	dsn, drv := drvFn(o)

	driver.
		sqlHooks := o.SQLHooks
	if len(sqlHooks) > 0 {
		drv = hooks.Wrap(drv, hooks.Compose(sqlHooks...))
	}
	var sqlDB *sql.DB
	conn := &dsnConnector{dsn: dsn, driver: drv}
	if o.UseOtel {
		sqlDB = otelsql.OpenDB(conn,
			otelsql.WithDBName(name),
			otelsql.WithDBSystem(dialect),
			otelsql.WithAttributes(attribute.String("ORM", "bun")),
		)
	} else {
		sqlDB = sql.OpenDB(conn)
	}
	// Ignore driver.ErrSkip when the Conn does not implement driver.Pinger interface
	if err := sqlDB.Ping(); err != nil && !errors.Is(err, driver.ErrSkip) {
		return nil, nil, err
	}

	connInitSQL := o.ConnInitSQL
	validationSQL := o.ValidationSQL
	if len(connInitSQL) > 0 {
		if _, err := sqlDB.Exec(connInitSQL); err != nil {
			return nil, nil, fmt.Errorf("unable to exec conn_init_sql: %s, reason: %s", connInitSQL, err)
		}
	}
	// Execute validation SQL after bun.DB is initialized
	if len(validationSQL) > 0 {
		if _, err := sqlDB.Exec(validationSQL); err != nil {
			return nil, nil, fmt.Errorf("unable to exec validation_sql: %s, reason: %s", validationSQL, err)
		}
	}

	// https://bun.uptrace.dev/guide/running-bun-in-production.html
	maxIdleConns := o.MaxIdleConns
	maxOpenConns := o.MaxOpenConns
	connMaxLifetime := o.ConnMaxLifetime
	connMaxIdleTime := o.ConnMaxIdleTime
	if maxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		// TODO In Docker how we should do?
		maxOpenConns = 4 * runtime.GOMAXPROCS(0)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	}
	if maxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(maxIdleConns)
	} else {
		sqlDB.SetMaxIdleConns(maxOpenConns)
	}
	if connMaxIdleTime != nil {
		sqlDB.SetConnMaxIdleTime(*connMaxIdleTime)
	}
	if connMaxLifetime != nil {
		sqlDB.SetConnMaxLifetime(*connMaxLifetime)
	}

	return dlt, sqlDB, err
}
