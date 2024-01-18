package sql

import (
	"context"
	"database/sql/driver"
)

var _ driver.Connector = (*dsnConnector)(nil)

type dsnConnector struct {
	dsn    string
	driver driver.Driver
}

func (t *dsnConnector) Connect(_ context.Context) (driver.Conn, error) {
	return t.driver.Open(t.dsn)
}

func (t *dsnConnector) Driver() driver.Driver {
	return t.driver
}
