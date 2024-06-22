package sqlite

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	xbun "github.com/blink-io/x/bun"
	xbunx "github.com/blink-io/x/bun/x"
	"github.com/google/uuid"
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

func TestSqlite_Bun_Update_Custom_1(t *testing.T) {
	db := getSqliteDB()

	rdb := xbunx.NewDB[Application, string](db)

	ds := rdb.NewUpdate().
		Table("applications").
		SetColumn("status", "?", "no-ok").
		Where("status = ?", "ok")
	_, err1 := ds.Exec(ctx)
	require.NoError(t, err1)
}

func TestSqlite_Bun_Update_Map_1(t *testing.T) {
	db := getSqliteDB()

	value := map[string]interface{}{
		"status":      "no-ok",
		"description": "updated-by-map",
	}
	_, err := db.NewUpdate().
		Model(&value).
		TableExpr("applications").
		Where("id = ?", 20).
		Exec(ctx)
	require.NoError(t, err)
}

func TestSqlite_Bun_Update_RawSQL_1(t *testing.T) {
	db := getSqliteDB()

	_, err := db.NewUpdate().NewRaw(
		"update applications set status = ?, description = ? where id = ?",
		"raw-sql-ok", "update-by-raw-sql", 22,
	).Exec(ctx)
	require.NoError(t, err)
}

func TestSqlite_Bun_Map_Insert_1(t *testing.T) {
	db := getSqliteDB()

	values := map[string]any{
		"guid":       uuid.NewString(),
		"username":   gofakeit.Name(),
		"location":   gofakeit.City(),
		"level":      gofakeit.Int8(),
		"profile":    gofakeit.AppName(),
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	sqlstr := db.NewInsert().
		Model(&values).
		Ignore().
		TableExpr("users").
		String()
	fmt.Println("sql: ", sqlstr)

	_, err := db.Exec(sqlstr)
	require.NoError(t, err)
}

func TestSqlite_Bun_Select_Custom_1(t *testing.T) {
	db := getSqliteDB()

	var ids []int64
	var descs []sql.Null[string]
	err := db.NewSelect().
		Table("applications").
		Column("id", "description").
		Limit(5).
		Scan(ctx, &ids, &descs)
	require.NoError(t, err)
}

func TestSqlite_Bun_Select_RawSQL_1(t *testing.T) {
	db := getSqliteDB()

	var users xbunx.ModelSlice[User]
	err := db.NewRaw("select * from users").Scan(ctx, &users)
	require.NoError(t, err)
}

func TestSqlite_Bun_All_Custom_1(t *testing.T) {
	db := getSqliteDB()
	ms, err := xbunx.All[IDAndName](ctx, db, xbunx.WithSelectQuery(func(q *xbun.SelectQuery) *xbun.SelectQuery {
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
