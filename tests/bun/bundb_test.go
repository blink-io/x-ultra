package bun

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/blink-io/x/sql"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func getDB(t *testing.T) *sql.DB {
	dbPath := filepath.Join(".", "bun_demo.db")

	fmt.Println("db path: ", dbPath)

	db, err1 := sql.NewDB(&sql.Options{
		Dialect: sql.DialectSQLite,
		Host:    dbPath,
	})
	require.NoError(t, err1)

	return db
}

func TestDB_SQLite_1(t *testing.T) {
	db := getDB(t)
	m := (*Application)(nil)
	_, err := db.NewCreateTable().IfNotExists().Model(m).Exec(ctx)
	require.NoError(t, err)
}
