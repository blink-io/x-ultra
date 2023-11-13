package hooks

import (
	"database/sql"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setUpPostgres(t *testing.T, dsn string) {
	db, err := sql.Open("pgx", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())
	defer db.Close()

	_, err = db.Exec("CREATE table IF NOT EXISTS users(id int, name text)")
	require.NoError(t, err)
}

func TestPostgres(t *testing.T) {
	dsn := os.Getenv("SQLHOOKS_POSTGRES_DSN")
	if dsn == "" {
		t.Skipf("SQLHOOKS_POSTGRES_DSN not set")
	}

	setUpPostgres(t, dsn)

	s := newSuite(t, stdlib.GetDefaultDriver(), dsn)

	s.TestHooksExecution(t, "SELECT * FROM users WHERE id = $1", 1)
	s.TestHooksArguments(t, "SELECT * FROM users WHERE id = $1 AND name = $2", int64(1), "Gus")
	s.TestHooksErrors(t, "SELECT 1+1")
	s.TestErrHookHook(t, "SELECT * FROM users WHERE id = $2", "INVALID_ARGS")

	t.Run("DBWorks", func(t *testing.T) {
		s.hooks.noop()
		if _, err := s.db.Exec("DELETE FROM users"); err != nil {
			t.Fatal(err)
		}

		stmt, err := s.db.Prepare("INSERT INTO users (id, name) VALUES($1, $2)")
		require.NoError(t, err)
		for i := range [5]struct{}{} {
			_, err := stmt.Exec(i, "gus")
			require.NoError(t, err)
		}

		var count int
		require.NoError(t,
			s.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count),
		)
		assert.Equal(t, 5, count)
	})
}
