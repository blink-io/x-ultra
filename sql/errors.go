package sql

import (
	"database/sql"
	"errors"

	//"github.com/glebarez/go-sqlite"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
)

type ErrState string

const (
	ErrStateUnsupported ErrState = "unsupported"

	ErrStateOther ErrState = "other"

	ErrStateNoRows ErrState = "now_rows"

	ErrStateTooManyRows ErrState = "too_many_rows"

	ErrStateConstraintUnique ErrState = "unique_constraint"

	ErrStateConstraintCheck ErrState = "check_constraint"

	ErrStateConstraintNotNull ErrState = "not_null_constraint"

	ErrStateConstraintForeignKey ErrState = "foreign_key_constraint"
)

func (s ErrState) String() string {
	return string(s)
}

// ToError creates *StateError with only state value is assigned.
func (s ErrState) ToError() *StateError {
	return &StateError{state: s}
}

func (s ErrState) NewError(code string, message string, cause error) *StateError {
	return NewStateError(s, code, message, cause)
}

var _ error = (*StateError)(nil)

var (
	ErrOther = ErrStateOther.ToError()

	ErrUnsupported = ErrStateUnsupported.ToError()

	ErrNoRows = ErrStateUnsupported.ToError()

	ErrTooManyRows = ErrStateTooManyRows.ToError()

	ErrConstraintUnique = ErrStateConstraintUnique.ToError()

	ErrConstraintCheck = ErrStateConstraintCheck.ToError()

	ErrConstraintNotNull = ErrStateConstraintNotNull.ToError()

	ErrConstraintForeignKey = ErrStateConstraintForeignKey.ToError()
)

type StateError struct {
	cause error

	// state defines unique id for error
	state ErrState

	// code in PostgreSQL/SQLite, number in MySQL
	code string

	message string
}

func (e *StateError) Error() string {
	return e.message + " (SQLSTATE " + string(e.state) + ")"
}

func (e *StateError) State() ErrState {
	return e.state
}

func (e *StateError) Code() string {
	return e.code
}

func (e *StateError) Cause() error {
	return e.cause
}

// Is when target is *StateError and their states are the same.
func (e *StateError) Is(target error) bool {
	return IsErrEqualsState(target, e.state)
}

func (e *StateError) Clone() *StateError {
	return NewStateError(e.state, e.code, e.message, e.cause)
}

func (e *StateError) Renew(code string, message string, cause error) *StateError {
	return NewStateError(e.state, code, message, cause)
}

func NewStateError(state ErrState, code string, message string, cause error) *StateError {
	return &StateError{
		state:   state,
		code:    code,
		message: message,
		cause:   cause,
	}
}

// WrapError wraps *pgconn.PgError/*mysql.MySQLError/sqlite3.Error to *StateError.
func WrapError(e error) *StateError {
	var newErr *StateError
	if sErr := new(StateError); errors.As(e, &sErr) {
		newErr = sErr
	} else if pgErr := new(pgconn.PgError); errors.As(e, &pgErr) {
		newErr = pgxStateError(pgErr)
	} else if mysqlErr := new(mysql.MySQLError); errors.As(e, &mysqlErr) {
		newErr = mysqlStateError(mysqlErr)
	} else if sqliteErr := new(SQLiteError); errors.As(e, sqliteErr) {
		newErr = sqliteStateError(sqliteErr)
	} else if ch, ok := commonErrHandlers[e]; ok {
		newErr = ch(e)
	} else {
		newErr = ErrUnsupported
	}
	return newErr
}

func IsErrNoRows(e error) bool {
	if errors.Is(e, sql.ErrNoRows) {
		return true
	}
	return IsErrEqualsState(e, ErrStateNoRows)
}

func IsErrTooManyRows(e error) bool {
	return IsErrEqualsState(e, ErrStateTooManyRows)
}

func IsErrConstraintCheck(e error) bool {
	return IsErrEqualsState(e, ErrStateConstraintCheck)
}

func IsErrConstraintUnique(e error) bool {
	return IsErrEqualsState(e, ErrStateConstraintUnique)
}

func IsErrConstraintNotNull(e error) bool {
	return IsErrEqualsState(e, ErrStateConstraintNotNull)
}

func IsErrConstraintForeignKey(e error) bool {
	return IsErrEqualsState(e, ErrStateConstraintForeignKey)
}

func IsErrEqualsState(e error, state ErrState) bool {
	if se := new(StateError); errors.As(e, &se) {
		return se.state == state
	}
	if we := WrapError(e); !errors.Is(we, ErrUnsupported) {
		return we.state == state
	}
	return false
}
