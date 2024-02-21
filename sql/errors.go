package sql

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
)

type ErrName string

const (
	ErrNameUnsupported ErrName = "unsupported"

	ErrNameOther ErrName = "other"

	ErrNameNoRows ErrName = "now_rows"

	ErrNameTooManyRows ErrName = "too_many_rows"

	ErrNameConstraintUnique ErrName = "unique_constraint"

	ErrNameConstraintCheck ErrName = "check_constraint"

	ErrNameConstraintNotNull ErrName = "not_null_constraint"

	ErrNameConstraintForeignKey ErrName = "foreign_key_constraint"
)

func (s ErrName) String() string {
	return string(s)
}

// ToError creates *Error with only name value is assigned.
func (s ErrName) ToError() *Error {
	return &Error{name: s}
}

func (s ErrName) NewError(code string, message string, cause error) *Error {
	return NewError(s, code, message, cause)
}

var _ error = (*Error)(nil)

var (
	ErrOther = ErrNameOther.ToError()

	ErrUnsupported = ErrNameUnsupported.ToError()

	ErrNoRows = ErrNameUnsupported.ToError()

	ErrTooManyRows = ErrNameTooManyRows.ToError()

	ErrConstraintUnique = ErrNameConstraintUnique.ToError()

	ErrConstraintCheck = ErrNameConstraintCheck.ToError()

	ErrConstraintNotNull = ErrNameConstraintNotNull.ToError()

	ErrConstraintForeignKey = ErrNameConstraintForeignKey.ToError()
)

type Error struct {
	cause error

	// name defines unique id for error
	name ErrName

	// code in PostgreSQL/SQLite, number in MySQL
	code string

	message string
}

func (e *Error) Error() string {
	return e.message + " (ERR_NAME " + string(e.name) + ")"
}

func (e *Error) Name() ErrName {
	return e.name
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Cause() error {
	return e.cause
}

// Is when target is *Error and their names are the same.
func (e *Error) Is(target error) bool {
	return IsErrEquals(target, e.name)
}

func (e *Error) Clone() *Error {
	return NewError(e.name, e.code, e.message, e.cause)
}

func (e *Error) Renew(code string, message string, cause error) *Error {
	return NewError(e.name, code, message, cause)
}

func NewError(name ErrName, code string, message string, cause error) *Error {
	return &Error{
		name:    name,
		code:    code,
		message: message,
		cause:   cause,
	}
}

// WrapError wraps *pgconn.PgError/*mysql.MySQLError/sqlite3.Error to *Error.
func WrapError(e error) *Error {
	var newErr *Error
	if tErr, ok := isTargetErr[*Error](e); ok {
		newErr = tErr
	} else if tErr, ok := isTargetErr[*pgconn.PgError](e); ok {
		newErr = handlePgxError(tErr)
	} else if tErr, ok := isTargetErr[*mysql.MySQLError](e); ok {
		newErr = handleMysqlError(tErr)
	} else if tErr, ok := isTargetErr[*SQLiteError](e); ok {
		newErr = handleSqliteError(tErr)
	} else if ef, ok := handleCommonError(e); ok {
		newErr = ef(e)
	} else {
		newErr = ErrUnsupported
	}
	return newErr
}

func isTargetErr[T error](e error) (T, bool) {
	tErr := new(T)
	ok := errors.As(e, tErr)
	return *tErr, ok
}

func IsErrNoRows(e error) bool {
	if errors.Is(e, sql.ErrNoRows) {
		return true
	}
	return IsErrEquals(e, ErrNameNoRows)
}

func IsErrTooManyRows(e error) bool {
	return IsErrEquals(e, ErrNameTooManyRows)
}

func IsErrConstraintCheck(e error) bool {
	return IsErrEquals(e, ErrNameConstraintCheck)
}

func IsErrConstraintUnique(e error) bool {
	return IsErrEquals(e, ErrNameConstraintUnique)
}

func IsErrConstraintNotNull(e error) bool {
	return IsErrEquals(e, ErrNameConstraintNotNull)
}

func IsErrConstraintForeignKey(e error) bool {
	return IsErrEquals(e, ErrNameConstraintForeignKey)
}

func IsErrEquals(e error, name ErrName) bool {
	if se, ok := isTargetErr[*Error](e); ok {
		return se.name == name
	}
	if we := WrapError(e); !errors.Is(we, ErrUnsupported) {
		return we.name == name
	}
	return false
}
