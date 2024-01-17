package sqlite

import (
	"fmt"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

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

func TestSqlite_DBR_Insert_1(t *testing.T) {
	db := getSqliteDBR()

	r1 := newRandomRecordForApp("dbr")
	r2 := newRandomRecordForApp("dbr")
	_, err := db.InsertInto(r1.Table()).
		Columns(r1.Columns()...).
		Record(r1).Record(r2).
		Exec()
	require.NoError(t, err)
}

func TestSqlite_DBR_Select_1(t *testing.T) {
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
