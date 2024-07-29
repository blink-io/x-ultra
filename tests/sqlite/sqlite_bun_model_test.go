package sqlite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/sanity-io/litter"

	bunx "github.com/blink-io/x/bun"
	"github.com/stretchr/testify/require"
)

func TestSqlite_Bun_Model_Update_2(t *testing.T) {
	db := getSqliteDB()
	xdb := bunx.NewGenericDB[Application, int64](db)

	r1, err := xdb.One(ctx)
	require.NoError(t, err)

	r1.Type = gofakeit.Name() + "-Updated"

	err = xdb.Update(ctx, r1)
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Delete_1(t *testing.T) {
	db := getSqliteDB()
	gdb := bunx.NewGenericDB[Application, int64](db)
	err := gdb.BulkDelete(ctx, []int64{5})
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Delete_Bulk_1(t *testing.T) {
	db := getSqliteDB()
	//
	//u1 := new(User)
	//u1.I = 1
	//
	//u2 := new(User)
	//u2.I = 2
	//
	//users := xbunx.ModelSlice[User]{u1, u2}
	xdb := bunx.NewGenericDB[User, int64](db)

	err := xdb.BulkDelete(ctx, []int64{1, 2})
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Delete_All(t *testing.T) {
	db := getSqliteDB()

	xdb := bunx.NewGenericDB[Application, int64](db)

	err := xdb.Delete(ctx, 16)
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Select_All_HasMany(t *testing.T) {
	db := getSqliteDB()

	records, err := bunx.DoAll[UserWithDevices](ctx, db)
	require.NoError(t, err)
	require.NotNil(t, records)
}

func TestSqlite_Bun_Model_BulkInsert_1(t *testing.T) {
	db := getSqliteDB()

	rrLen := 10
	rr := make([]*Application, rrLen)
	for i := 0; i < rrLen; i++ {
		rr[i] = newRandomRecordForApp(bunx.Accessor)
	}

	tdb, err := bunx.NewGenericDB[Application, int64](db).Tx(ctx, nil)
	require.NoError(t, err)

	err1 := tdb.BulkInsert(ctx, rr)
	require.NoError(t, err1)

	require.NoError(t, tdb.Commit())
}

func TestSqlite_Bun_Model_Insert_1(t *testing.T) {
	db := getSqliteDB()
	rrLen := 10
	rr := make([]*Application, rrLen)
	for i := 0; i < rrLen; i++ {
		r1 := newRandomRecordForApp(bunx.Accessor)
		rr[i] = r1
	}

	xdb := bunx.NewGenericDB[Application, int64](db)

	err1 := xdb.BulkInsert(ctx, rr, bunx.DoWithInsertReturning("id"))
	require.NoError(t, err1)
}

func TestSqlite_Bun_Model_UserDevice_Insert_1(t *testing.T) {
	db := getSqliteDB()
	rrLen := 10
	rr := make([]*UserDevice, rrLen)
	for i := 0; i < rrLen; i++ {
		r1 := newRandomRecordForUserDevice(bunx.Accessor)
		rr[i] = r1
	}

	xdb := bunx.NewGenericDB[UserDevice, int64](db)

	err1 := xdb.BulkInsert(ctx, rr, bunx.DoWithInsertReturning("id"))
	require.NoError(t, err1)
}

func TestSqlite_Bun_Model_One_1(t *testing.T) {
	db := getSqliteDB()
	m, err := bunx.DoOne[Application](ctx, db,
		bunx.DoWithSelectWhere("id = ?", 16),
	)

	require.NoError(t, err)
	require.NotNil(t, m)
}

func TestSqlite_Bun_Model_All_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := bunx.DoAll[Application](ctx, db,
		bunx.DoWithSelectQuery(func(q *bunx.SelectQuery) *bunx.SelectQuery {
			q.Limit(3)
			return q
		}),
	)
	require.NoError(t, err)
	require.NotNil(t, ms)

	fmt.Println("Is Empty: ", ms.Emtpy())
}

func TestSqlite_Bun_Model_Select_1(t *testing.T) {
	db := getSqliteDB()

	rt := new(Application)

	err := db.NewSelect().Model(rt).
		Where("type like ?", "%type-001%").
		Scan(ctx)
	require.NoError(t, err)
	litter.Dump(rt)
}

func TestSqlite_Bun_Model_Select_2(t *testing.T) {
	db := getSqliteDB()
	var rs []*Application
	_, err := db.NewRaw("select * from applications where ? = ? limit 3", bunx.Ident("status"), "ok").
		Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Select_Custom_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := bunx.Struct[IDAndProfile](ctx, db, bunx.DoWithSelectQuery(func(q *bunx.SelectQuery) *bunx.SelectQuery {
		q.Table("users")
		q.Limit(3)
		return q
	}))
	require.NoError(t, err)
	require.NotNil(t, ms)
}

func TestSqlite_Bun_Model_Select_Custom_2(t *testing.T) {
	db := getSqliteDB()
	ms, err := bunx.StructSQL[IDAndProfile](ctx, db, "select * from users limit ?", 10)
	require.NoError(t, err)
	require.NotNil(t, ms)
}

func TestSqlite_Bun_Model_Count_1(t *testing.T) {
	db := getSqliteDB()
	c, err := bunx.DoCount[Application](ctx, db)
	require.NoError(t, err)
	fmt.Println("DoCount: ", c)
}

func TestSqlite_Bun_Model_Exists_1(t *testing.T) {
	db := getSqliteDB()
	e1, err := bunx.DoExists[Application](ctx, db,
		bunx.DoWithSelectQuery(func(q *bunx.SelectQuery) *bunx.SelectQuery {
			q.Where("id = ?", 18)
			return q
		}),
	)
	require.NoError(t, err)
	assert.Equal(t, e1, true)

	e2, err := bunx.DoExists[Application](ctx, db,
		bunx.DoWithSelectQuery(func(q *bunx.SelectQuery) *bunx.SelectQuery {
			q.Where("id = ?", 118)
			return q
		}),
	)
	require.NoError(t, err)
	assert.Equal(t, e2, false)
}
