//go:build cgo

package sql

import (
	"database/sql/driver"

	"github.com/mattn/go-sqlite3"
)

func getRawSQLiteDriver() driver.Driver {
	return &sqlite3.SQLiteDriver{}
}
