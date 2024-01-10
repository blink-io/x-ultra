package sql

import (
	"context"
	"database/sql/driver"
	"time"
)

type (
	Dsner = func(context.Context, *Config) (string, error)

	dialectOptions struct {
		loc *time.Location
	}

	// DialectOption defines option for dialect
	DialectOption func(*dialectOptions)

	GetDriverFunc func(dialect string) driver.Driver

	GetDSNFunc func(dialect string) (Dsner, error)
)

var (
	drivers = make(map[string]GetDriverFunc)

	dsners = make(map[string]GetDSNFunc)
)
