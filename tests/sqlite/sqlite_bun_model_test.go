package sqlite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/sanity-io/litter"

	xbun "github.com/blink-io/x/bun"
	xbunx "github.com/blink-io/x/bun/x"
	"github.com/stretchr/testify/require"
)

func TestSqlite_Bun_Model_Update_2(t *testing.T) {
	db := getSqliteDB()
	xdb := xbunx.NewDB[Application, int64](db)

	r1, err := xdb.One(ctx)
	require.NoError(t, err)

	r1.Type = gofakeit.Name() + "-Updated"

	err = xdb.Update(ctx, r1)
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Delete_1(t *testing.T) {
	db := getSqliteDB()
	gdb := xbunx.NewDB[Application, int64](db)
	err := gdb.BulkDelete(ctx, []int64{5})
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Delete_All(t *testing.T) {
	db := getSqliteDB()

	xdb := xbunx.NewDB[Application, int64](db)

	err := xdb.Delete(ctx, 16)
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_BulkInsert_1(t *testing.T) {
	db := getSqliteDB()

	rrLen := 10
	rr := make([]*Application, rrLen)
	for i := 0; i < rrLen; i++ {
		rr[i] = newRandomRecordForApp(xbun.Accessor)
	}

	tdb, err := xbunx.NewDB[Application, int64](db).Tx(ctx, nil)
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
		r1 := newRandomRecordForApp(xbun.Accessor)
		rr[i] = r1
	}

	xdb := xbunx.NewDB[Application, int64](db)

	err1 := xdb.BulkInsert(ctx, rr, xbunx.WithInsertReturning("id"))
	require.NoError(t, err1)
}

func TestSqlite_Bun_Model_One_1(t *testing.T) {
	db := getSqliteDB()
	m, err := xbunx.One[Application](ctx, db,
		xbunx.WithSelectWhere("id = ?", 16),
	)

	require.NoError(t, err)
	require.NotNil(t, m)
}

func TestSqlite_Bun_Model_All_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := xbunx.All[Application](ctx, db,
		xbunx.WithSelectQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
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
	_, err := db.NewRaw("select * from applications where ? = ? limit 3", xbun.Ident("status"), "ok").
		Exec(ctx, &rs)
	require.NoError(t, err)
}

func TestSqlite_Bun_Model_Count_1(t *testing.T) {
	db := getSqliteDB()
	c, err := xbunx.Count[Application](ctx, db)
	require.NoError(t, err)
	fmt.Println("Count: ", c)
}

func TestSqlite_Bun_Model_Exists_1(t *testing.T) {
	db := getSqliteDB()
	e1, err := xbunx.Exists[Application](ctx, db,
		xbunx.WithSelectQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
			q.Where("id = ?", 18)
			return q
		}),
	)
	require.NoError(t, err)
	assert.Equal(t, e1, true)

	e2, err := xbunx.Exists[Application](ctx, db,
		xbunx.WithSelectQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
			q.Where("id = ?", 118)
			return q
		}),
	)
	require.NoError(t, err)
	assert.Equal(t, e2, false)
}
