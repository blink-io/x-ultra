package sql

import (
	"github.com/jackc/pgx/v5/pgconn"
)

func pgxStateErr(e *pgconn.PgError) *StateError {
	err := &StateError{
		cause:   e,
		code:    e.Code,
		message: e.Message,
	}
	return err
}
