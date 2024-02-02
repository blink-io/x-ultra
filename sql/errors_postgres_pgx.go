package sql

import (
	"github.com/jackc/pgx/v5/pgconn"
)

// pgxStateErr transform PgError to StateError
// Doc: https://www.postgresql.org/docs/11/protocol-error-fields.html
func pgxStateErr(e *pgconn.PgError) *StateError {
	err := &StateError{
		cause:   e,
		code:    e.Code,
		message: e.Message,
	}
	return err
}

func IsErrPgxUniqueContraint(e *pgconn.PgError) bool {
	return false
}
