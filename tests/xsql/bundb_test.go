package bun

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/sql/generics"
	"github.com/blink-io/x/sql/scany/dbscan"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func TestDB_SQLite_1(t *testing.T) {
	db := getDBWithSQLite()
	m := (*Application)(nil)
	_, err := db.NewCreateTable().IfNotExists().Model(m).Exec(ctx)
	require.NoError(t, err)
}

func TestSQLite3_Select_Funcs(t *testing.T) {
	db := getDBWithSQLite()

	sqlF := "select %s as payload"
	funcs := []string{
		"hex(randomblob(32))",
		"random()",
		"sqlite_version()",
		"total_changes()",
		`lower("HELLO")`,
		`upper("hello")`,
		`length("hello")`,
		`length("我是世界")`,
		//`concat("Hello", ",", "World")`,
	}

	for _, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		row := db.QueryRow(ss)
		var v string
		require.NoError(t, row.Scan(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}

func TestSQLite3_Delete_1(t *testing.T) {
	db := getDBWithSQLite()
	gdb := generics.NewDB[Application, string](db)
	//err := gdb.Delete(ctx, "123456")
	err := gdb.BulkDelete(ctx, []string{"123456", "888888"})
	require.NoError(t, err)
}

func TestSQLite3_Insert_1(t *testing.T) {
	db := getDBWithSQLite()
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
	db := getDBWithSQLite()

	qm := &Application{
		Status: "status3",
	}
	_, err := db.NewSelect().Model(qm).
		Where("status = ?", "status3").
		Exec(ctx)
	require.NoError(t, err)
}

func TestSQLite_Raw_Select_Model_1(t *testing.T) {
	db := getDBWithSQLite()
	var rs []*Application
	_, err := db.NewRaw("select * from applications where ? = ?",
		bun.Ident("status"), "status1").Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSQLite_Raw_Select_Model_2(t *testing.T) {
	db := getDBWithSQLite()
	var rs []Application
	_, err := db.NewRaw("select * from applications where ? = ?",
		bun.Ident("status"), "status1").Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSQLite_Raw_Select_2(t *testing.T) {
	db := getDBWithSQLite()

	defer db.Close()

	_, err := db.Exec("select * from applications where status = ?", "status1")
	require.NoError(t, err)
}

func TestSLite_Scan_Slice_1(t *testing.T) {
	db := getDBWithSQLite()

	defer db.Close()

	var rs []Application

	rows, err := db.Query("select * from applications limit 5")
	require.NoError(t, err)
	err = dbscan.ScanAll(&rs, rows)
	require.NoError(t, err)
}

func TestSLite_Scan_Slice_2(t *testing.T) {
	db := getDBWithSQLite()

	defer db.Close()

	var rs []*Application

	rows, err := db.Query("select * from applications limit 5")
	require.NoError(t, err)
	err = dbscan.ScanAll(&rs, rows)
	require.NoError(t, err)
}

func TestSLite_Scan_Slice_3(t *testing.T) {
	db := getDBWithSQLite()

	defer db.Close()

	var rs []map[string]any

	rows, err := db.Query("select * from applications limit 5")
	require.NoError(t, err)
	err = dbscan.ScanAll(&rs, rows)
	require.NoError(t, err)
}

func TestSLite_Scan_2(t *testing.T) {
	db := getDBWithSQLite()

	defer db.Close()

	var rs = new(Application)

	rows, err := db.Query("select * from applications limit 1")
	require.NoError(t, err)
	err = dbscan.ScanOne(rs, rows)
	require.NoError(t, err)
}
