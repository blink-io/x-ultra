package sqlite

import (
	"fmt"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
	"github.com/vingarcia/ksql"
)

func TestSqlite_DBK_Select_Funcs(t *testing.T) {
	db := getSqliteDBK()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	type Res struct {
		Payload string `ksql:"payload"`
	}

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		var rt Res
		err := db.QueryOne(ctx, &rt, ss)
		require.NoError(t, err)
		fmt.Println(k, "======>", rt.Payload)
	}
}

func TestSqlite_DBK_Insert_1(t *testing.T) {
	db := getSqliteDBK()

	var AppTable = ksql.NewTable("applications")

	r1 := newRandomRecordForApp("dbr")
	r2 := newRandomRecordForApp("dbr")

	err1 := db.Insert(ctx, AppTable, r1)
	require.NoError(t, err1)

	err2 := db.Insert(ctx, AppTable, r2)
	require.NoError(t, err2)
}

func TestSqlite_DBK_Select_1(t *testing.T) {
	db := getSqliteDBR()

	rt := new(Application)

	err := db.Select("*").From(rt.Table()).
		Where("type like ?", "%type-001%").
		Where("? = ?", 1, 1).
		Limit(1).LoadOne(rt)
	require.NoError(t, err)
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}
