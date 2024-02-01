//go:build !cgo

package sql

import (
	"github.com/blink-io/x/cast"

	"modernc.org/sqlite"
)

type SQLiteError = sqlite.Error

func sqliteStateErr(e *SQLiteError) *StateError {
	err := &StateError{
		cause:   e,
		code:    cast.ToString(e.Code()),
		message: e.Error(),
	}
	return err
}
