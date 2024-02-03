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
	"github.com/google/uuid"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBQ_Dialect(t *testing.T) {
	db := getSqliteDBQ()
	require.NotNil(t, db)
	sql := "select * from applications limit 10;"

	var as []*Application

	err := db.ScanStructs(&as, sql)
	require.NoError(t, err)
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

func TestSqlite_DBQ_Tx_Insert_Succ_1(t *testing.T) {
	db := getSqliteDBQ()
	tx, err := db.Begin()
	require.NoError(t, err)

	r1 := newRandomRecordForApp("dbq")
	r2 := newRandomRecordForApp("dbq")

	ds := tx.From(r1.TableName())
	_, err = ds.Insert().Rows(r1, r2).
		Executor().Exec()

	if err != nil {
		require.NoError(t, tx.Rollback())
		fmt.Println("invoke tx rollback")
	} else {
		require.NoError(t, tx.Commit())
		fmt.Println("invoke tx commit")
	}
}

func TestSqlite_DBQ_Tx_Insert_Fail_1(t *testing.T) {
	db := getSqliteDBQ()
	tx, err := db.Begin()
	require.NoError(t, err)

	r1 := newRandomRecordForApp("dbq")
	r2 := newRandomRecordForApp("dbq")

	sameCode := id.ShortID()
	r1.Code = sameCode
	r2.Code = sameCode

	ds := tx.From(r1.TableName())
	_, err = ds.Insert().Rows(r1, r2).
		Executor().Exec()

	if err != nil {
		fmt.Println("Err: ", err)
		require.NoError(t, tx.Rollback())
		fmt.Println("invoke tx rollback")
	} else {
		require.NoError(t, tx.Commit())
		fmt.Println("invoke tx commit")
	}
}

func TestSqlite_DBQ_Update_1(t *testing.T) {
	db := getSqliteDBQ()

	ds := db.Update("applications").Set(
		goqu.Record{"status": "no-ok"},
	).Where(goqu.C("id").Gt(8))
	updateSQL, args, _ := ds.ToSQL()
	fmt.Println(updateSQL, args)

	var ra = new(Application)

	ok, err := ds.Returning(goqu.Star()).Executor().ScanStruct(ra)
	require.NoError(t, err)

	fmt.Println("OK? ", ok)

	litter.Dump(ra)
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

func TestSqlite_DBQ_Insert_NoDS(t *testing.T) {
	db := getSqliteDBQ()
	require.NotNil(t, db)

	//mm := appModel()
	//
	//r1 := newRandomRecordForApp("dbq")
	//r2 := newRandomRecordForApp("dbq"
	//insertSQL, args, _ := ds.Insert().Rows(r1).ToSQL()
	//fmt.Println(insertSQL, args)
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
		Cols("id", "guid", "type", "status", "code", "name", "created_at", "updated_at", "deleted_at").
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

func TestSqlite_DBQ_Insert_3(t *testing.T) {
	n := time.Now().Local()
	db := getSqliteDBQ()
	sql, args, err := db.Insert("applications").Prepared(true).
		//Cols("id", "iid", "type", "status", "code", "name", "created_at", "updated_at", "deleted_at").
		Rows(
			goqu.Record{
				"guid":        uuid.NewString(),
				"type":        "type-003",
				"code":        id.ShortUUID(),
				"name":        gofakeit.Name(),
				"status":      "very-ok",
				"description": gofakeit.Company(),
				"created_at":  n,
				"updated_at":  n,
			},
			goqu.Record{
				"guid":        uuid.NewString(),
				"type":        "type-003",
				"code":        id.ShortUUID(),
				"name":        gofakeit.Name(),
				"status":      "very-bad",
				"description": gofakeit.Company(),
				"created_at":  n,
				"updated_at":  n,
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

	insertSQL, args, _ := ds.Insert().Rows(r1).Prepared(true).ToSQL()
	fmt.Println(insertSQL, args)
}

func TestSqlite_DBQ_Delete_All(t *testing.T) {
	db := getSqliteDBQ()

	_, err := db.Delete("applications").Executor().Exec()
	require.NoError(t, err)
}

func TestSqlite_DBQ_Delete_Returning_1(t *testing.T) {
	db := getSqliteDBQ()

	var da = new(Application)

	ok, err := db.Delete("applications").
		Where(goqu.Ex{"id": 8}).
		Returning(goqu.Star()).Executor().
		ScanStruct(da)
	require.NoError(t, err)

	fmt.Println("OK? ", ok)
}
