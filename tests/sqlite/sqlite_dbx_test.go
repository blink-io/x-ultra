package sqlite

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBX_Select_Funcs(t *testing.T) {
	db := getSqliteDBX()

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

func TestSqlite_DBX_Select_NoRows(t *testing.T) {
	db := getSqliteDBX()
	require.NotNil(t, db)

}

func TestSqlite_DBX_WrapError_NoRows(t *testing.T) {
	db := getSqliteDBX()
	sql := "select * from users where id = 18876"

	rows, err := db.QueryContext(ctx, sql)
	require.NoError(t, err)
	require.NotNil(t, rows)

	var user User
	errv := dbscan.ScanOne(&user, rows)
	require.NoError(t, errv)
}
