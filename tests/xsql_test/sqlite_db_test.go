package xsql_test

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/sql/generics"
	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/go-rel/rel"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func TestSqlite_DBX_Select_Funcs(t *testing.T) {
	db := getSqliteDBX()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	for _, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		q := db.NewQuery(ss)
		var v string
		require.NoError(t, q.Row(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}

func TestSqlite_DBX_Insert_1(t *testing.T) {
	db := getSqliteDBX()

	r1 := newRandomRecordForApp("dbx")

	err := db.Model(r1).Insert()
	require.NoError(t, err)
}

func TestSqlite_DBR_Select_1(t *testing.T) {
	db := getSqliteDBR()

	sess := db.NewSession(nil)

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for _, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		var v string
		r := sess.QueryRow(ss)
		require.NoError(t, r.Scan(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}

func TestSqlite_DBQ_Select_Funcs(t *testing.T) {
	//Debug = true
	db := getSqliteDBQ()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		row := db.QueryRow(ss)
		var v string
		require.NoError(t, row.Scan(&v))
		fmt.Println(k, ":  ", v)
	}
}

func TestSqlite_DBQ_Insert_1(t *testing.T) {
	db := getSqliteDBQ()

	r1 := newRandomRecordForApp("goqu")
	ds := db.From(r1.TableName())
	//_, err := ds.Insert().Rows(r1).
	//	Executor().Exec()
	//require.NoError(t, err)
	insertSQL, args, _ := ds.Insert().Rows(r1).ToSQL()
	fmt.Println(insertSQL, args)
}

func TestSqlite_DB_CreateTable_1(t *testing.T) {
	db := getSqliteDB()
	m := (*Application)(nil)
	_, err := db.NewCreateTable().IfNotExists().Model(m).Exec(ctx)
	require.NoError(t, err)
}

func TestSqlite_DB_CreateTable_Funcs(t *testing.T) {
	db := getSqliteDB()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for _, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		row := db.QueryRow(ss)
		var v string
		require.NoError(t, row.Scan(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}

func TestSqlite_DB_Delete_1(t *testing.T) {
	db := getSqliteDB()
	gdb := generics.NewDB[Application, string](db)
	//err := gdb.Delete(ctx, "123456")
	err := gdb.BulkDelete(ctx, []string{"123456", "888888"})
	require.NoError(t, err)
}

func TestSqlite_DB_BulkInsert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp("bun")
	r2 := newRandomRecordForApp("bun")
	r3 := newRandomRecordForApp("bun")

	tdb, err := generics.NewDB[Application, string](db).Tx()
	require.NoError(t, err)

	err1 := tdb.BulkInsert(ctx, []*Application{r1, r2, r3})
	require.NoError(t, err1)

	require.NoError(t, tdb.Commit())
}

func TestSqlite_DB_SelectModel_1(t *testing.T) {
	db := getSqliteDB()

	qm := &Application{
		Status: "status3",
	}
	_, err := db.NewSelect().Model(qm).
		Where("status = ?", "status3").
		Exec(ctx)
	require.NoError(t, err)
}

func TestSqlite_DB_SelectModel_2(t *testing.T) {
	db := getSqliteDB()
	var rs []*Application
	_, err := db.NewRaw("select * from applications where ? = ?",
		bun.Ident("status"), "status1").Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSqlite_DBScan_1(t *testing.T) {
	db := getSqliteDB()

	defer db.Close()

	var rs []map[string]any

	rows, err := db.Query("select * from applications limit 5")
	require.NoError(t, err)
	err = dbscan.ScanAll(&rs, rows)
	require.NoError(t, err)
}

func TestSqlite_DBScan_2(t *testing.T) {
	db := getSqliteDB()

	defer db.Close()

	var rs = new(Application)

	rows, err := db.Query("select * from applications limit 1")
	require.NoError(t, err)
	err = dbscan.ScanOne(rs, rows)
	require.NoError(t, err)
}

func TestSqlite_DBM_Select_Funcs(t *testing.T) {
	db := getSqliteDBM()

	type Result struct {
		Payload string `db:"payload"`
	}

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	rt := new(Result)
	for k, v := range funcs {
		ss := rel.SQL(fmt.Sprintf(sqlF, v))
		require.NoError(t, db.Find(ctx, rt, ss))
		fmt.Println(k, "=>", rt.Payload)
	}
}

func TestSqlite_DBM_Select_1(t *testing.T) {
	db := getSqliteDBM()

	type Result struct {
		Verinfo  string `db:"verinfo"`
		Rstr     string `db:"rstr"`
		SourceID string `db:"source_id"`
	}

	var rtmap Result

	sqlstr := `select
random() as rstr,
sqlite_source_id() as source_id, 
sqlite_version() as verinfo;`

	err := db.Find(ctx, &rtmap, rel.SQL(sqlstr))
	require.NoError(t, err)

	fmt.Printf("------------------------------------------------------------\n")

	fmt.Println("Sqlite version: ", rtmap)

	require.NoError(t, db.Ping(ctx))
}
