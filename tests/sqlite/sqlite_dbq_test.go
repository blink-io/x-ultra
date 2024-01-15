package sqlite

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sqlite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/doug-martin/goqu/v9"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

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

	sqlF := "select * from applications where type like ? limit 1"

	var v Application

	_, err := db.ScanStruct(&v, sqlF, "%type-01%")
	require.NoError(t, err)
	fmt.Println("Record: ", v)
}

func TestSqlite_DBQ_Select_2(t *testing.T) {
	//Debug = true
	db := getSqliteDBQ()

	ds := db.From("applications")

	var rt = new(Application)

	sel := ds.Select(ToAnySlice[string](appColumns)...).Where(goqu.L("type like ?", "%type-01%"))

	_, err := sel.ScanStruct(rt)
	require.NoError(t, err)
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}

func TestSqlite_DBQ_Select_All(t *testing.T) {
	//Debug = true
	db := getSqliteDBQ()

	ds := db.From("applications")

	var rts []*Application

	err1 := ds.Select(ToAnySlice[string](appColumns)...).
		Where(goqu.L("type = ?", "type-001")).ScanStructs(&rts)

	require.NoError(t, err1)
	fmt.Println("Result ==================>", len(rts))
}

func TestSqlite_DBQ_Select_3(t *testing.T) {
	sql, _, _ := goqu.Select(goqu.L("NOW()")).ToSQL()
	fmt.Println(sql)
}

func TestSqlite_DBQ_Update_1(t *testing.T) {
	db := getSqliteDBQ()

	ds := db.Update("applications").Set(
		goqu.Record{"status": "no-ok"},
	).Where(goqu.Ex{
		"type":   "type-005",
		"status": "ok",
	})
	updateSQL, args, _ := ds.ToSQL()
	fmt.Println(updateSQL, args)
}

func TestSqlite_DBQ_Update_2(t *testing.T) {
	ds := goqu.From("user")

	updateSQL, _, _ := ds.Update().Set(
		goqu.Record{"first_name": "Greg", "last_name": "Farley"},
	).ToSQL()
	fmt.Println(updateSQL)

	updateSQL, _, _ = ds.Where(goqu.C("first_name").Eq("Gregory")).Update().Set(
		goqu.Record{"first_name": "Greg", "last_name": "Farley"},
	).ToSQL()
	fmt.Println(updateSQL)

	r1 := newRandomRecordForApp("dbq")
	sql, args, _ := goqu.Update("applications").Set(
		r1,
	).ToSQL()
	fmt.Println(sql, args)
}

func TestSqlite_DBQ_Select_Version(t *testing.T) {
	db := getSqliteDBQ()

	ver := sqlite.QueryVersion(ctx, db.QueryRowContext)
	fmt.Println("Version: ", ver)
}

func TestSqlite_DBQ_Insert_1(t *testing.T) {
	db := getSqliteDBQ()

	r1 := newRandomRecordForApp("dbq")
	r2 := newRandomRecordForApp("dbq")
	ds := db.From(r1.TableName())
	_, err := ds.Insert().Rows(r1, r2).
		Executor().Exec()
	require.NoError(t, err)
	//insertSQL, args, _ := ds.Insert().Rows(r1).ToSQL()
	//fmt.Println(insertSQL, args)
}

func TestSqlite_DBQ_Insert_2(t *testing.T) {
	n := time.Now().Local()
	db := getSqliteDBQ()
	sql, args, err := db.Insert("applications").Prepared(true).
		Cols("id", "iid", "type", "status", "code", "name", "created_at", "updated_at", "deleted_at").
		Vals(
			goqu.Vals{
				id.GUID(),
				int64(gofakeit.Int32()),
				"type-005",
				"ok",
				"code-" + gofakeit.Name(),
				"name-" + gofakeit.Name(),
				n,
				n,
				xsql.ValidTime(n),
			},
		).ToSQL()
	require.NoError(t, err)
	fmt.Println("sql: ", sql, ", args: ", args)

	_, err2 := db.ExecContext(ctx, sql, args...)
	require.NoError(t, err2)
}

func TestSqlite_DBQ_Insert_SQLGen(t *testing.T) {
	db := getSqliteDBQ()

	r1 := newRandomRecordForApp("dbq")
	ds := db.From(r1.TableName())

	insertSQL, args, _ := ds.Insert().Rows(r1).ToSQL()
	fmt.Println(insertSQL, args)
}
