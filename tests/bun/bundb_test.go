package bun

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/blink-io/x/sql"
	"github.com/google/uuid"
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

func TestSQLite3_Select_Version(t *testing.T) {
	ss := "select sqlite_version() as version"
	db := getDB(t)
	row := db.QueryRow(ss)
	var v string
	require.NoError(t, row.Scan(&v))

	fmt.Println("SQLite version:  ", v)
}

func TestSQLite3_Insert_1(t *testing.T) {
	db := getDB(t)
	r1 := &Application{}
	r1.ID = uuid.New().String()
	r1.Name = "app1"
	r1.Code = "code1"
	r1.Type = "type1"
	r1.Status = "status1"
	_, err := db.NewInsert().Model(r1).Exec(context.Background())
	require.NoError(t, err)
}
