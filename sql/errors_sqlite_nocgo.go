//go:build !cgo

package sql

import (
	"github.com/blink-io/x/cast"
	"modernc.org/sqlite"
)

type SQLiteError = *sqlite.Error

var sqliteErrorHandlers = map[int]func(*sqlite.Error) *StateError{
	// SQLITE_CONSTRAINT_CHECK (275)
	275: func(e *sqlite.Error) *StateError {
		return ErrConstraintCheck.
			Renew(cast.ToString(e.Code()), e.Error(), e)
	},
	// SQLITE_CONSTRAINT_FOREIGNKEY (787)
	787: func(e *sqlite.Error) *StateError {
		return ErrConstraintForeignKey.
			Renew(cast.ToString(e.Code()), e.Error(), e)
	},
	// SQLITE_CONSTRAINT_NOTNULL (1299)
	1299: func(e *sqlite.Error) *StateError {
		return ErrConstraintNotNull.
			Renew(cast.ToString(e.Code()), e.Error(), e)
	},
	// SQLITE_CONSTRAINT_PRIMARYKEY (1555).
	// Notes: In DBMS, primary key is a unique key too.
	1555: sqliteUniqueConstraintHandler,
	// SQLITE_CONSTRAINT_UNIQUE (2067)
	2067: sqliteUniqueConstraintHandler,
}

func RegisterSQLiteErrorHandler(number int, fn func(*sqlite.Error) *StateError) {
	sqliteErrorHandlers[number] = fn
}

func sqliteUniqueConstraintHandler(e *sqlite.Error) *StateError {
	return ErrConstraintUnique.
		Renew(cast.ToString(e.Code()), e.Error(), e)
}

func sqliteStateError(e *sqlite.Error) *StateError {
	if h, ok := sqliteErrorHandlers[e.Code()]; ok {
		return h(e)
	} else {
		return ErrOther.Renew(cast.ToString(e.Code()), e.Error(), e)
	}
}
