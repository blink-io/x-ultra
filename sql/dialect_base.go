package sql

import (
	"context"
	"database/sql/driver"
)

type (
	Dsner = func(context.Context, *Config) (string, error)

	GetDriverFunc func(dialect string) (driver.Driver, error)

	GetDSNFunc func(dialect string) (Dsner, error)
)

var (
	drivers = make(map[string]GetDriverFunc)

	dsners = make(map[string]GetDSNFunc)
)
