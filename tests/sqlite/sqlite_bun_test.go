package sqlite

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/stretchr/testify/assert"

	xbun "github.com/blink-io/x/bun"
	xbunx "github.com/blink-io/x/bun/x"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

var m1 = (*Application)(nil)
var m2 = (*User)(nil)
var ms = []any{m1, m2}

func TestSqlite_Bun_CreateTable_1(t *testing.T) {
	db := getSqliteDB()
	for _, m := range ms {
		_, err := db.NewCreateTable().IfNotExists().Model(m).Exec(ctx)
		require.NoError(t, err)
	}
}

func TestSqlite_Bun_DropTable_1(t *testing.T) {
	db := getSqliteDB()
	for _, m := range ms {
		_, err := db.NewDropTable().IfExists().Model(m).Exec(ctx)
		require.NoError(t, err)
	}
}

func TestSqlite_Bun_RebuildTable_1(t *testing.T) {
	TestSqlite_Bun_DropTable_1(t)
	TestSqlite_Bun_CreateTable_1(t)
}

func TestSqlite_Bun_Delete_1(t *testing.T) {
	db := getSqliteDB()
	gdb := xbunx.NewDB[Application, int64](db)
	//err := gdb.Delete(ctx, "123456")
	err := gdb.BulkDelete(ctx, []int64{5})
	require.NoError(t, err)
}

func TestSqlite_Bun_Insert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp(xbun.Accessor)

	rdb := xbunx.NewDB[Application, string](db)

	err1 := rdb.Insert(ctx, r1, xbunx.WithInsertReturning("id"))
	require.NoError(t, err1)
}

func TestSqlite_Bun_BulkInsert_1(t *testing.T) {
	db := getSqliteDB()
	r1 := newRandomRecordForApp(xbun.Accessor)
	r2 := newRandomRecordForApp(xbun.Accessor)
	r3 := newRandomRecordForApp(xbun.Accessor)

	tdb, err := xbunx.NewDB[Application, string](db).Tx(ctx, nil)
	require.NoError(t, err)

	err1 := tdb.BulkInsert(ctx, []*Application{r1, r2, r3})
	require.NoError(t, err1)

	require.NoError(t, tdb.Commit())
}

func TestSqlite_Bun_Update_1(t *testing.T) {
	db := getSqliteDB()

	rdb := xbunx.NewDB[Application, string](db)

	ds := rdb.NewUpdate().
		Table("applications").
		SetColumn("status", "?", "no-ok").
		Where("status = ?", "ok")
	_, err1 := ds.Exec(ctx)
	require.NoError(t, err1)
}

func TestSqlite_Bun_Update_2(t *testing.T) {
	db := getSqliteDB()
	gdb := xbunx.NewDB[Application, int64](db)

	r1, err := gdb.Get(ctx, 6)
	require.NoError(t, err)

	r1.Name = gofakeit.Name() + "-Updated"

	err = gdb.Update(ctx, r1)
	require.NoError(t, err)
}

func TestSqlite_Bun_Delete_All(t *testing.T) {
	db := getSqliteDB()

	rdb := xbunx.NewDB[Application, string](db)

	ds := rdb.NewDelete()

	_, err1 := ds.Exec(ctx)
	require.NoError(t, err1)
}

func TestSqlite_Bun_InsertMap_1(t *testing.T) {
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

func TestSqlite_Bun_SelectCols_1(t *testing.T) {
	db := getSqliteDB()

	var ids []int64
	var desc []*string
	err := db.NewSelect().
		Table("applications").
		Column("id", "description").
		Limit(5).
		Scan(ctx, &ids, &desc)
	require.NoError(t, err)
}

func TestSqlite_Bun_SelectTypeTuple2_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple2[int64, string](ctx, db,
		"applications", "id", "description", xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple3_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple3[int64, string, string](ctx, db,
		"applications", "id", "name", "description", xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple4_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple4[int64, string, string, string](ctx, db,
		"applications", "id", "name", "code", "description", xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple5_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple5[int64, string, string, string, string](ctx, db,
		"applications", "id", "name", "status", "code", "description", xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple6_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple6[int64, sql.Null[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple7_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple7[int64, string, sql.Null[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "guid", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple8_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple8[int64, int64, string, sql.Null[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "id", "guid", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple9_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple9[int64, int64, string, string, sql.Null[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "id", "guid", "guid", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
		xbunx.WithSelectWhere("id > ? and level is not null", 0),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectModel_1(t *testing.T) {
	db := getSqliteDB()

	rt := new(Application)

	err := db.NewSelect().Model(rt).
		Where("type like ?", "%type-01%").
		Scan(ctx)
	require.NoError(t, err)
	litter.Dump(rt)
}

func TestSqlite_Bun_SelectModel_2(t *testing.T) {
	db := getSqliteDB()
	var rs []*Application
	_, err := db.NewRaw("select * from applications where ? = ?",
		xbun.Ident("status"), "status1").Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSqlite_Bun_One_1(t *testing.T) {
	db := getSqliteDB()
	m, err := xbunx.One[Application](ctx, db, xbunx.WithSelectWhere("id = ?", 6))

	require.NoError(t, err)
	require.NotNil(t, m)
}

func TestSqlite_Bun_Get_1(t *testing.T) {
	db := getSqliteDB()
	guid := "318620d3-199b-4613-ad7c-762af7ae43a0"
	m, err := xbunx.Get[Application](ctx, db, guid, "guid")

	require.NoError(t, err)
	require.NotNil(t, m)
	assert.Equal(t, guid, m.GUID)
}

func TestSqlite_Bun_All_Model_1(t *testing.T) {
	db := getSqliteDB()
	//var rs xbunx.ModelSlice[Application]
	ms, err := xbunx.All[Application](ctx, db, xbunx.WithSelectApplyQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
		q.Limit(3)
		return q
	}))
	require.NoError(t, err)
	require.NotNil(t, ms)

	println("Is Empty: ", ms.Emtpy())
}

func TestSqlite_Bun_Type_All_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := xbunx.Type[string](ctx, db, "applications", "guid",
		xbunx.WithSelectApplyQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
			q.Limit(3)
			return q
		}))
	require.NoError(t, err)
	require.NotNil(t, ms)

	fmt.Println(ms)
}

func TestSqlite_Bun_All_Custom_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := xbunx.All[IDAndName](ctx, db, xbunx.WithSelectApplyQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
		q.ModelTableExpr("applications as a1")
		q.Limit(3)
		return q
	}))
	require.NoError(t, err)
	require.NotNil(t, ms)

	println("Is Empty: ", ms.Emtpy())
}

func TestSqlite_Bun_Select_Funcs(t *testing.T) {
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
