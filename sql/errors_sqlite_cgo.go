//go:build cgo

package sql

import (
	"github.com/blink-io/x/cast"
	"github.com/mattn/go-sqlite3"
)

type SQLiteError = sqlite3.Error

var sqliteErrorHandlers = map[int]func(*sqlite3.Error) *StateError{
	// SQLITE_CONSTRAINT_CHECK (275)
	275: func(e *sqlite3.Error) *StateError {
		code := cast.ToString(int(e.ExtendedCode))
		return ErrConstraintCheck.Renew(code, e.Error(), e)
	},
	// SQLITE_CONSTRAINT_FOREIGNKEY (787)
	787: func(e *sqlite3.Error) *StateError {
		code := cast.ToString(int(e.ExtendedCode))
		return ErrConstraintForeignKey.Renew(code, e.Error(), e)
	},
	// SQLITE_CONSTRAINT_NOTNULL (1299)
	1299: func(e *sqlite3.Error) *StateError {
		code := cast.ToString(int(e.ExtendedCode))
		return ErrConstraintNotNull.Renew(code, e.Error(), e)
	},
	// SQLITE_CONSTRAINT_PRIMARYKEY (1555).
	// Notes: In DBMS, primary key is a unique key too.
	1555: sqliteUniqueConstraintHandler,

	// SQLITE_CONSTRAINT_UNIQUE (2067)
	2067: sqliteUniqueConstraintHandler,
}

func RegisterSQLiteErrorHandler(number int, fn func(*sqlite3.Error) *StateError) {
	sqliteErrorHandlers[number] = fn
}

func sqliteUniqueConstraintHandler(e *sqlite3.Error) *StateError {
	code := cast.ToString(int(e.ExtendedCode))
	return ErrConstraintUnique.Renew(code, e.Error(), e)
}

// sqliteStateError transforms *sqlite3.Error to *StateError.
// Doc: https://www.sqlite.org/rescode.html
func sqliteStateError(e *sqlite3.Error) *StateError {
	code := int(e.ExtendedCode)
	if h, ok := sqliteErrorHandlers[code]; ok {
		return h(e)
	} else {
		return ErrOther.Renew(cast.ToString(code), e.Error(), e)
	}
}
