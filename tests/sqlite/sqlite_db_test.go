package sqlite

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/blink-io/x/sql/db/g"
	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/blink-io/x/sqlite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/huandu/go-sqlbuilder"
	"github.com/sanity-io/litter"
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

func TestSqlite_DBX_CreateTable_Model8(t *testing.T) {

}

func TestSqlite_DBX_Select_All(t *testing.T) {
	db := getSqliteDBX()

	var as []*Application
	err := db.Select().From("applications").All(&as)
	require.NoError(t, err)

	fmt.Println("Result count: ", len(as))
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

func TestSqlite_DBR_Select_1(t *testing.T) {
	db := getSqliteDBR()

	rt := new(Application)

	_, err := db.Select("*").From(rt.Table()).
		Where("type like ?", "%type-01%").
		Where("? = ?", 1, 1).Load(rt)
	require.NoError(t, err)
	fmt.Println("Result ==================>")
	litter.Dump(rt)
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
	gdb := g.NewDB[Application, string](db)
	//err := gdb.Delete(ctx, "123456")
	err := gdb.BulkDelete(ctx, []string{"123456", "888888"})
	require.NoError(t, err)
}

func TestSqlite_DB_Insert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp(xdb.Accessor)

	rdb := g.NewDB[Application, string](db)

	err1 := rdb.Insert(ctx, r1, g.InsertReturning("id"))
	require.NoError(t, err1)
}

func TestSqlite_DB_BulkInsert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp(xdb.Accessor)
	r2 := newRandomRecordForApp(xdb.Accessor)
	r3 := newRandomRecordForApp(xdb.Accessor)

	tdb, err := g.NewDB[Application, string](db).Tx()
	require.NoError(t, err)

	err1 := tdb.BulkInsert(ctx, []*Application{r1, r2, r3})
	require.NoError(t, err1)

	require.NoError(t, tdb.Commit())
}

func TestSqlite_DB_Update_1(t *testing.T) {
	db := getSqliteDB()

	rdb := g.NewDB[Application, string](db)

	ds := rdb.NewUpdate().
		Table("applications").
		SetColumn("status", "?", "no-ok").
		Where("status = ?", "ok")
	_, err1 := ds.Exec(ctx)
	require.NoError(t, err1)
}

func TestSqlite_DB_Delete_All(t *testing.T) {
	db := getSqliteDB()

	rdb := g.NewDB[Application, string](db)

	ds := rdb.NewDelete()

	_, err1 := ds.Exec(ctx)
	require.NoError(t, err1)
}

func TestSqlite_DB_InsertMap_1(t *testing.T) {
	db := getSqliteDB()

	values := map[string]interface{}{
		"title": "title1",
		"text":  "text1",
	}
	sql := db.NewInsert().
		Model(&values).
		Ignore().
		TableExpr("books").
		String()
	fmt.Println("sql: ", sql)
}

func TestSqlite_DB_SelectModel_1(t *testing.T) {
	db := getSqliteDB()

	rt := new(Application)

	err := db.NewSelect().Model(rt).
		Where("type like ?", "%type-01%").
		Scan(ctx)
	require.NoError(t, err)
	litter.Dump(rt)
}

func TestSqlite_DB_SelectModel_2(t *testing.T) {
	db := getSqliteDB()
	var rs []*Application
	_, err := db.NewRaw("select * from applications where ? = ?",
		bun.Ident("status"), "status1").Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSqlite_DB_Select_Funcs(t *testing.T) {
	db := getSqliteDB()

	type Result struct {
		Payload string `db:"payload"`
	}

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		row := db.QueryRowContext(ctx, ss)
		var rstr string
		err := row.Scan(&rstr)
		require.NoError(t, err)
		fmt.Println(k, "=>", rstr)
	}
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

func TestSqlite_DBM_Select_2(t *testing.T) {
	db := getSqliteDBM()
	rt := new(Application)

	sqlF := "select * from applications where type like ? limit 1"
	ss := rel.SQL(sqlF, "%type-01%")
	require.NoError(t, db.Find(ctx, rt, ss))
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}

func TestSqlite_DBM_Select_3(t *testing.T) {
	db := getSqliteDBM()
	rt := new(Application)

	require.NoError(t, db.Find(ctx, rt, where.Like("type", "%type-01%")))
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}

func TestSqlite_DBM_Select_Version(t *testing.T) {
	db := getSqliteDBM().SqlDB()
	ver := sqlite.QueryVersion(ctx, db.QueryRowContext)
	fmt.Println("Version: ", ver)
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

	sqlF := "select * from applications where type like ? limit 1"

	err1 := db.SelectOne(rt, sqlF, "%type-01%")
	require.NoError(t, err1)
	fmt.Println("Result ======================> ")
	litter.Dump(rt)
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

	_, err := db.Exec(sql, args...)
	require.NoError(t, err)
}

func TestSqlite_DBW_Select_Funcs(t *testing.T) {
	db := getSqliteDBW()

	sqlF := "select %s as payload where ? = ?"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		var rt string
		err1 := db.Get(&rt, ss, 1, 1)
		require.NoError(t, err1)
		fmt.Println(k, "=>", rt)
	}
}

func getXSqlOpts(d, n string) *xsql.Config {
	opts := &xsql.Config{
		Dialect: d,
		Name:    n,
	}
	return opts
}

func TestSqlOptions(t *testing.T) {
	opts1 := getXSqlOpts("ddd", "nnnn")
	opts2 := getXSqlOpts("ddd", "nnnn")

	require.NotNil(t, opts1)
	require.NotNil(t, opts2)
}

func TestSqlite_GoSqlBuilder_1(t *testing.T) {
	var appStruct = sqlbuilder.NewStruct(new(Application))

	sb := appStruct.SelectFrom("applications")
	sb.Where(sb.Equal("id", "GY7hPxdSg"))

	// Execute the query.
	sql, args := sb.Build()
	//rows, _ := db.Query(sql, args...)

	fmt.Println("SQL: ", sql)
	fmt.Println("Args: ", args)

	db := getSqliteSqlDB()
	rows, err := db.Query(sql, args...)
	require.NoError(t, err)

	var rs []*Application
	require.NoError(t, dbscan.ScanAll(&rs, rows))

	fmt.Println("done")
}
