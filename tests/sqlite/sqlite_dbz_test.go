package sqlite

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBZ_Select_Funcs(t *testing.T) {
	db := getSqliteDBZ()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		rows, err := db.QueryContext(ctx, ss)
		require.NoError(t, err)

		var str string
		err1 := dbscan.ScanOne(&str, rows)
		require.NoError(t, err1)

		fmt.Println(k, "----->", str)
	}
}

func TestSqlite_DBZ_Select_NoRows(t *testing.T) {
	db := getSqliteDBZ()
	require.NotNil(t, db)

}

func TestSqlite_DBZ_Select_All(t *testing.T) {
	db := getSqliteDBX()

	var as []*Application
	err := db.Select().From("applications").All(&as)
	require.NoError(t, err)

	fmt.Println("Result count: ", len(as))
}

func TestSqlite_DBZ_Insert_1(t *testing.T) {
	db := getSqliteDBX()

	r1 := newRandomRecordForApp("dbx")

	err := db.Model(r1).Insert()
	require.NoError(t, err)
}

func TestSqlite_DBZ_WrapError_NoRows(t *testing.T) {
	db := getSqliteDBZ()
	sql := "select * from users where id = 18876"

	rows, err := db.QueryContext(ctx, sql)
	require.NoError(t, err)
	require.NotNil(t, rows)

	var user User
	errv := dbscan.ScanOne(&user, rows)
	require.NoError(t, errv)
}
