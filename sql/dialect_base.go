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
)

var (
	drivers = make(map[string]driver.Driver)

	dsners = make(map[string]Dsner)
)
