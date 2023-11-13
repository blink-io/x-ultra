package hooks

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	"modernc.org/sqlite"
)

func init() {
	hooks := &testHooks{}
	hooks.noop()

	sql.Register("sqlite3-benchmark", Wrap(&sqlite.Driver{}, hooks))
	sql.Register("mysql-benchmark", Wrap(&mysql.MySQLDriver{}, hooks))
	sql.Register("postgres-benchmark", Wrap(stdlib.GetDefaultDriver(), hooks))
}

func benchmark(b *testing.B, driver, dsn string) {
	db, err := sql.Open(driver, dsn)
	require.NoError(b, err)
	defer db.Close()

	var query = "SELECT 'hello'"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, err := db.Query(query)
		require.NoError(b, err)
		require.NoError(b, rows.Close())
	}
}

func BenchmarkSQLite3(b *testing.B) {
	b.Run("Without Hooks", func(b *testing.B) {
		benchmark(b, "sqlite", ":memory:")
	})

	b.Run("With Hooks", func(b *testing.B) {
		benchmark(b, "sqlite3-benchmark", ":memory:")
	})
}

func BenchmarkMySQL(b *testing.B) {
	dsn := os.Getenv("SQLHOOKS_MYSQL_DSN")
	if dsn == "" {
		b.Skipf("SQLHOOKS_MYSQL_DSN not set")
	}

	b.Run("Without Hooks", func(b *testing.B) {
		benchmark(b, "mysql", dsn)
	})

	b.Run("With Hooks", func(b *testing.B) {
		benchmark(b, "mysql-benchmark", dsn)
	})
}

func BenchmarkPostgres(b *testing.B) {
	dsn := os.Getenv("SQLHOOKS_POSTGRES_DSN")
	if dsn == "" {
		b.Skipf("SQLHOOKS_POSTGRES_DSN not set")
	}

	b.Run("Without Hooks", func(b *testing.B) {
		benchmark(b, "postgres", dsn)
	})

	b.Run("With Hooks", func(b *testing.B) {
		benchmark(b, "postgres-benchmark", dsn)
	})
}
