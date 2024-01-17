package sql

import (
	"errors"

	"github.com/blink-io/x/cast"
	//"github.com/glebarez/go-sqlite"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
	"modernc.org/sqlite"
)

type StateError struct {
	origin  error
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

func WrapError(e error) *StateError {
	newErr := &StateError{
		origin: e,
	}
	if pgErr := new(pgconn.PgError); errors.As(e, &pgErr) {
		newErr.code = pgErr.Code
		newErr.message = pgErr.Message
	} else if mysqlErr := new(mysql.MySQLError); errors.As(e, &mysqlErr) {
		newErr.code = cast.ToString(mysqlErr.Number)
		newErr.message = mysqlErr.Message
	} else if sqliteErr := new(sqlite.Error); errors.As(e, &sqliteErr) {
		newErr.code = cast.ToString(sqliteErr.Code())
		newErr.message = sqliteErr.Error()
	} else {
		newErr.message = e.Error()
	}
	return newErr
}
