package sqlite

import (
	"fmt"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

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
