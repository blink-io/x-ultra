package sql

import (
	"github.com/lib/pq"
)

var pqErrorHandlers = map[string]func(*pq.Error) *Error{
	// P0002	no_data_found
	"P0002": func(e *pq.Error) *Error {
		return ErrNoRows.Renew(string(e.Code), e.Message, e)
	},
	// Class 02 — No Data (this is also a warning class per the SQL standard)
	// 02000	no_data
	"02000": func(e *pq.Error) *Error {
		return ErrNoRows.Renew(string(e.Code), e.Message, e)
	},
	// P0003	too_many_rows
	"P0003": func(e *pq.Error) *Error {
		return ErrTooManyRows.Renew(string(e.Code), e.Message, e)
	},
	// 23502	not_null_violation
	"23502": func(e *pq.Error) *Error {
		return ErrConstraintNotNull.Renew(string(e.Code), e.Message, e)
	},
	// 23503	foreign_key_violation
	"23503": func(e *pq.Error) *Error {
		return ErrConstraintForeignKey.Renew(string(e.Code), e.Message, e)
	},
	// 23505	unique_violation
	"23505": func(e *pq.Error) *Error {
		return ErrConstraintUnique.Renew(string(e.Code), e.Message, e)
	},
	// 23514	check_violation
	"23514": func(e *pq.Error) *Error {
		return ErrConstraintCheck.Renew(string(e.Code), e.Message, e)
	},
}

func RegisterPqErrorHandler(code string, fn func(*pq.Error) *Error) {
	pqErrorHandlers[code] = fn
}

// pqError transforms *pq.Error to *Error.
// Doc: https://www.postgresql.org/docs/11/protocol-error-fields.html.
func pqError(e *pq.Error) *Error {
	if h, ok := pqErrorHandlers[string(e.Code)]; ok {
		return h(e)
	} else {
		return ErrOther.Renew(string(e.Code), e.Message, e)
	}
}
