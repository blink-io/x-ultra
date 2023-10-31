package bun

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/blink-io/x/bun/extra/logging"
	"github.com/blink-io/x/bun/extra/timing"
	"github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/generics"
	"github.com/blockloop/scan/v2"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
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
	db.AddQueryHook(logging.Func(log.Printf))
	db.AddQueryHook(timing.New())
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

func TestSQLite3_Delete_1(t *testing.T) {
	db := getDB(t)
	gdb := generics.NewDB[Application, string](db)
	//err := gdb.Delete(ctx, "123456")
	err := gdb.BulkDelete(ctx, []string{"123456", "888888"})
	require.NoError(t, err)
}

func TestSQLite3_Insert_1(t *testing.T) {
	db := getDB(t)
	r1 := &Application{}
	r1.ID = "123456"
	r1.Name = "app2"
	r1.Code = "code2"
	r1.Type = "type2"
	r1.Status = "status2"

	r3 := &Application{}
	r3.ID = "888888"
	r3.Name = "app3"
	r3.Code = "code3"
	r3.Type = "type3"
	r3.Status = "status3"

	tdb, err := generics.NewDB[Application, string](db).Tx()
	require.NoError(t, err)

	err1 := tdb.BulkInsert(ctx, []*Application{r1, r3})
	require.NoError(t, tdb.Commit())
	require.NoError(t, err1)
}

func TestSQLite_Model_Select_2(t *testing.T) {
	db := getDB(t)

	qm := &Application{
		Status: "status3",
	}
	_, err := db.NewSelect().Model(qm).
		Where("status = ?", "status3").
		Exec(ctx)
	require.NoError(t, err)
}

func TestSQLite_Raw_Select_1(t *testing.T) {
	db := getDB(t)
	var rs []*Application
	_, err := db.NewRaw("select * from applications where ? = ?",
		bun.Ident("status"), "status1").Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSQLite_Raw_Select_2(t *testing.T) {
	db := getDB(t)

	defer db.Close()

	_, err := db.Exec("select * from applications where status = ?", "status1")
	require.NoError(t, err)
}

func TestSLite_Scan_1(t *testing.T) {
	db := getDB(t)

	defer db.Close()

	var rs []Application

	rows, err := db.Query("select * from applications where  1=1")
	require.NoError(t, err)
	err = scan.Rows(&rs, rows)
	require.NoError(t, err)
}
