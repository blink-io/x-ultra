package sqlite

import (
	"fmt"
	"testing"

	xdb "github.com/blink-io/x/sql/db"
	xdbx "github.com/blink-io/x/sql/db/x"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

var m1 = (*Application)(nil)
var m2 = (*User)(nil)
var ms = []any{m1, m2}

func TestSqlite_DB_CreateTable_1(t *testing.T) {
	db := getSqliteDB()
	for _, m := range ms {
		_, err := db.NewCreateTable().IfNotExists().Model(m).Exec(ctx)
		require.NoError(t, err)
	}
}

func TestSqlite_DB_DropTable_1(t *testing.T) {
	db := getSqliteDB()
	for _, m := range ms {
		_, err := db.NewDropTable().IfExists().Model(m).Exec(ctx)
		require.NoError(t, err)
	}
}

func TestRebuildTable_1(t *testing.T) {
	TestSqlite_DB_DropTable_1(t)
	TestSqlite_DB_CreateTable_1(t)
}

func TestSqlite_DB_Delete_1(t *testing.T) {
	db := getSqliteDB()
	gdb := xdbx.NewDB[Application, string](db)
	//err := gdb.Delete(ctx, "123456")
	err := gdb.BulkDelete(ctx, []string{"123456", "888888"})
	require.NoError(t, err)
}

func TestSqlite_DB_Insert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp(xdb.Accessor)

	rdb := xdbx.NewDB[Application, string](db)

	err1 := rdb.Insert(ctx, r1, xdbx.InsertReturning("id"))
	require.NoError(t, err1)
}

func TestSqlite_DB_BulkInsert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp(xdb.Accessor)
	r2 := newRandomRecordForApp(xdb.Accessor)
	r3 := newRandomRecordForApp(xdb.Accessor)

	tdb, err := xdbx.NewDB[Application, string](db).Tx()
	require.NoError(t, err)

	err1 := tdb.BulkInsert(ctx, []*Application{r1, r2, r3})
	require.NoError(t, err1)

	require.NoError(t, tdb.Commit())
}

func TestSqlite_DB_Update_1(t *testing.T) {
	db := getSqliteDB()

	rdb := xdbx.NewDB[Application, string](db)

	ds := rdb.NewUpdate().
		Table("applications").
		SetColumn("status", "?", "no-ok").
		Where("status = ?", "ok")
	_, err1 := ds.Exec(ctx)
	require.NoError(t, err1)
}

func TestSqlite_DB_Delete_All(t *testing.T) {
	db := getSqliteDB()

	rdb := xdbx.NewDB[Application, string](db)

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
