package sql

import (
	"database/sql"
	"errors"

	//"github.com/glebarez/go-sqlite"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrCodeUndefined = "undefined"

	ErrCodeConstraintUnique = "unique_constraint"

	ErrCodeConstraintNotNull = "unique_constraint"
)

var (
	ErrConstraintPrimaryKey = &StateError{}
	ErrConstraintUnique     = &StateError{}
	ErrConstraintCheck      = &StateError{}
	ErrConstraintNotNull    = &StateError{}
)

type StateError struct {
	cause error

	message string
	// Code in Postgres/SQLite, Number in MySQl
	code string
}

func (e *StateError) Error() string {
	return e.message + " (SQLSTATE " + e.code + ")"
}

func (e *StateError) Code() string {
	return e.code
}

func (e *StateError) Cause() error {
	return e.cause
}

func (e *StateError) Is(err error) bool {
	return errors.Is(e.cause, err)
}

var E = WrapError

func WrapError(e error) *StateError {
	var newErr *StateError
	if pgErr := new(pgconn.PgError); errors.As(e, &pgErr) {
		newErr = pgxStateErr(pgErr)
	} else if mysqlErr := new(mysql.MySQLError); errors.As(e, &mysqlErr) {
		newErr = mysqlStateErr(mysqlErr)
	} else if sqliteErr := new(SQLiteError); errors.As(e, &sqliteErr) {
		newErr = sqliteStateErr(sqliteErr)
	} else {
		newErr = &StateError{
			cause:   e,
			code:    ErrCodeUndefined,
			message: e.Error(),
		}
	}
	return newErr
}

func IsNoRowsErr(e error) bool {
	return errors.Is(e, sql.ErrNoRows)
}
