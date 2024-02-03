package sql

import (
	"errors"
	"fmt"
	"testing"

	"github.com/blink-io/x/cast"
	berrors "github.com/blink-io/x/errors"

	"modernc.org/sqlite"
	//"github.com/glebarez/go-sqlite"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestPgErr(t *testing.T) {
	var err = &pgconn.PgError{Message: "PgErr", Code: DialectPostgres}

	newErr := WrapError(err)

	require.Equal(t, newErr.message, err.Message)
	require.Equal(t, newErr.code, err.Code)
}

func TestMySQLErr(t *testing.T) {
	var err = &mysql.MySQLError{Message: "MySQLErr", Number: 8888}

	newErr := WrapError(err)

	require.Equal(t, newErr.message, err.Message)
	require.Equal(t, newErr.code, cast.ToString(err.Number))
}

func TestSQLiteErr(t *testing.T) {
	var err = &sqlite.Error{}

	newErr := WrapError(err)

	require.NotNil(t, newErr)
}

func TestErrEqual(t *testing.T) {
	var cause1 = errors.New("cause1")
	var cause2 = berrors.New("cause2 from blink-x")
	var cause3 = errors.New("cause2 from blink-x")
	var err1 = NewStateError("good", "very good1", "", cause1)
	var err2 = NewStateError("good", "very good2", "", cause2)
	var err3 = err1.Renew("babamama", "Very BabaMama", cause3)

	b1 := errors.Is(err1, err2)
	b2 := errors.Is(err3, err2)
	require.True(t, b1)
	require.True(t, b2)
}

func TestStateError_Clone(t *testing.T) {
	var cause1 = errors.New("cause1")
	var err1 = NewStateError(ErrStateConstraintUnique, "very good1", "", cause1)
	var err2 = err1.Clone()
	var err3 = WrapError(&pgconn.PgError{Code: "23505"})
	var err4 = WrapError(&mysql.MySQLError{Number: uint16(1169)})

	b1 := errors.Is(err1, err2)
	b2 := errors.Is(err2, err3)
	b3 := errors.Is(err3, err4)
	b4 := errors.Is(err1, err4)

	fmt.Println("Result: ", b1)
	fmt.Println("Result: ", b2)
	fmt.Println("Result: ", b3)
	fmt.Println("Result: ", b4)
}
