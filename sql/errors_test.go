package sql

import (
	"testing"

	"github.com/blink-io/x/cast"

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
