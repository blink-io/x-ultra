package sql

import (
	"github.com/jackc/pgx/v5/pgconn"
)

var pgxErrorHandlers = map[string]func(*pgconn.PgError) *Error{
	// P0002	no_data_found
	"P0002": func(e *pgconn.PgError) *Error {
		return ErrNoRows.Renew(e.Code, e.Message, e)
	},
	// Class 02 — No Data (this is also a warning class per the SQL standard)
	// 02000	no_data
	"02000": func(e *pgconn.PgError) *Error {
		return ErrNoRows.Renew(e.Code, e.Message, e)
	},
	// P0003	too_many_rows
	"P0003": func(e *pgconn.PgError) *Error {
		return ErrTooManyRows.Renew(e.Code, e.Message, e)
	},
	// 23502	not_null_violation
	"23502": func(e *pgconn.PgError) *Error {
		return ErrConstraintNotNull.Renew(e.Code, e.Message, e)
	},
	// 23503	foreign_key_violation
	"23503": func(e *pgconn.PgError) *Error {
		return ErrConstraintForeignKey.Renew(e.Code, e.Message, e)
	},
	// 23505	unique_violation
	"23505": func(e *pgconn.PgError) *Error {
		return ErrConstraintUnique.Renew(e.Code, e.Message, e)
	},
	// 23514	check_violation
	"23514": func(e *pgconn.PgError) *Error {
		return ErrConstraintCheck.Renew(e.Code, e.Message, e)
	},
}

func RegisterPgxErrorHandler(code string, fn func(*pgconn.PgError) *Error) {
	pgxErrorHandlers[code] = fn
}

// handlePgxError transforms *pgconn.PgError to *Error.
// Doc: https://www.postgresql.org/docs/11/protocol-error-fields.html.
func handlePgxError(e *pgconn.PgError) *Error {
	if h, ok := pgxErrorHandlers[e.Code]; ok {
		return h(e)
	} else {
		return ErrOther.Renew(e.Code, e.Message, e)
	}
}
