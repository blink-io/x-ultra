package sqlite

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/aarondl/opt/omitnull"
	xbun "github.com/blink-io/x/bun"

	xbunx "github.com/blink-io/x/bun/x"
	"github.com/gofrs/uuid/v5"
	guuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSqlite_Bun_Type_All_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := xbunx.Type[string](ctx, db, "applications", "guid",
		xbunx.WithSelectQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
			q.Limit(3)
			return q
		}))
	require.NoError(t, err)
	require.NotNil(t, ms)

	fmt.Println(ms)
}

func TestSqlite_Bun_Type_All_2(t *testing.T) {
	db := getSqliteDB()
	ms, err := xbunx.Type[int64](ctx, db, "", "",
		xbunx.WithSelectQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
			q.Column("id")
			q.Table("applications")
			q.Limit(3)
			return q
		}))
	require.NoError(t, err)
	require.NotNil(t, ms)

	fmt.Println(ms)
}

func TestSqlite_Bun_SelectTypeTuple2_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple2[int64, string](ctx, db,
		"applications", "id", "", xbunx.WithSelectLimit(5),
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

	ts, err := xbunx.TypeTuple6[int64, omitnull.Val[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple7_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple7[int64, string, omitnull.Val[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "guid", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple8_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple8[int64, int64, guuid.UUID, omitnull.Val[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "id", "guid", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestSqlite_Bun_SelectTypeTuple9_1(t *testing.T) {
	db := getSqliteDB()

	ts, err := xbunx.TypeTuple9[int64, int64, guuid.UUID, uuid.UUID, omitnull.Val[int], string, string, string, sql.Null[string]](ctx, db,
		"applications", "id", "id", "guid", "guid", "level", "name", "status", "code", "description",
		xbunx.WithSelectLimit(5),
		xbunx.WithSelectWhere("id > ? and level is not null", 0),
	)
	require.NoError(t, err)
	require.NotNil(t, ts)
}
