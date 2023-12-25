package sqlite

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	"github.com/blink-io/x/sql/generics"
	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-rel/rel"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

func TestSqlite_DBX_Select_Funcs(t *testing.T) {
	db := getSqliteDBX()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		q := db.NewQuery(ss)
		var s string
		require.NoError(t, q.Row(&s))
		slog.Info("result: ", k, s)
	}
}

func TestSqlite_DBX_Insert_1(t *testing.T) {
	db := getSqliteDBX()

	r1 := newRandomRecordForApp("dbx")

	err := db.Model(r1).Insert()
	require.NoError(t, err)
}

func TestSqlite_DBR_Insert_1(t *testing.T) {
	db := getSqliteDBR()

	r1 := newRandomRecordForApp("dbr")
	_, err := db.InsertInto(r1.Table()).
		Columns(r1.Columns()...).
		Record(r1).
		Exec()
	require.NoError(t, err)
}

func TestSqlite_DBR_Select_Funcs(t *testing.T) {
	db := getSqliteDBR()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		var rt string
		r := db.QueryRow(ss)
		require.NoError(t, r.Scan(&rt))
		fmt.Println(k, "=>", rt)
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

func TestSqlite_DBQ_Select_1(t *testing.T) {
	//Debug = true
	db := getSqliteDBQ()

	sqlF := "select * from applications"

	var v Application

	_, err := db.ScanStruct(&v, sqlF)
	require.NoError(t, err)
	fmt.Println("Record: ", v)
}

func TestSqlite_DBQ_Insert_1(t *testing.T) {
	db := getSqliteDBQ()

	r1 := newRandomRecordForApp("dbq")
	ds := db.From(r1.TableName())
	_, err := ds.Insert().Rows(r1).
		Executor().Exec()
	require.NoError(t, err)
	//insertSQL, args, _ := ds.Insert().Rows(r1).ToSQL()
	//fmt.Println(insertSQL, args)
}

func TestSqlite_DBQ_Insert_SQLGen(t *testing.T) {
	db := getSqliteDBQ()

	r1 := newRandomRecordForApp("dbq")
	ds := db.From(r1.TableName())

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

func TestSqlite_DBM_Insert_1(t *testing.T) {
	db := getSqliteDBM()

	r1 := newRandomRecordForApp(db.Accessor())

	err := db.Insert(ctx, r1)
	require.NoError(t, err)
}

func TestSqlite_DBP_Insert_1(t *testing.T) {
	db := getSqliteDBP()
	db.AddTableWithName(Application{}, "applications").SetKeys(false, "ID")

	r1 := newRandomRecordForApp("dbp")

	err := db.Insert(r1)
	require.NoError(t, err)
}

func TestSqlite_DBP_Select_1(t *testing.T) {
	db := getSqliteDBP()

	var rt = new(Application)

	sqlF := "select * from applications limit 1"

	err1 := db.SelectOne(rt, sqlF)
	require.NoError(t, err1)
	fmt.Println("Result => ", rt)
}

func TestSqlite_DBP_Select_Funcs(t *testing.T) {
	db := getSqliteDBP()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		rt, err1 := db.SelectNullStr(ss)
		require.NoError(t, err1)
		fmt.Println(k, "=>", rt.String)
	}
}

func TestSqlite_DBW_Insert_1(t *testing.T) {
	db := getSqliteDBW()

	sql := "insert into applications (id, iid, name,status, code, type, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8)"

	args := []any{
		id.ShortUUID(),
		gofakeit.Int32(),
		"from-" + db.Accessor() + "-" + gofakeit.Name(),
		"ok",
		"001-" + id.ShortID(),
		"type-01",
		time.Now(),
		time.Now(),
	}
	//r1 := newRandomRecordForApp("dbp")

	_, err := db.ExecContext(ctx, sql, args...)
	require.NoError(t, err)
}

func TestSqlite_DBW_Select_Funcs(t *testing.T) {
	db := getSqliteDBW()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		var rt string
		err1 := db.GetContext(ctx, &rt, ss)
		require.NoError(t, err1)
		fmt.Println(k, "=>", rt)
	}
}