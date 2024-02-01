//go:build !cgo

package sql

import (
	"database/sql/driver"

	"modernc.org/sqlite"
)

func getRawSQLiteDriver() driver.Driver {
	return &sqlite.Driver{}
}
