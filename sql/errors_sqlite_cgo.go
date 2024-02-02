//go:build cgo

package sql

import (
	"github.com/blink-io/x/cast"

	"github.com/mattn/go-sqlite3"
)

type SQLiteError = sqlite3.Error

func sqliteStateErr(e *SQLiteError) *StateError {
	err := &StateError{
		cause:   e.Code,
		code:    cast.ToString(e.Code),
		message: e.Error(),
	}
	return err
}
