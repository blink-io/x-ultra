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

func (c *dsnConnector) Connect(_ context.Context) (driver.Conn, error) {
	return c.driver.Open(c.dsn)
}

func (c *dsnConnector) Driver() driver.Driver {
	return c.driver
}
