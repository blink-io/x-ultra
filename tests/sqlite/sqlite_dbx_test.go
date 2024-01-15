package sqlite

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqlite_DBX_Select_Funcs(t *testing.T) {
	db := getSqliteDBX()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		q := db.NewQuery(ss)
		var s string
		require.NoError(t, q.Row(&s))
		slog.Info("result: ", k, s)
	}
}

func TestSqlite_DBX_CreateTable_Model8(t *testing.T) {

}

func TestSqlite_DBX_Select_All(t *testing.T) {
	db := getSqliteDBX()

	var as []*Application
	err := db.Select().From("applications").All(&as)
	require.NoError(t, err)

	fmt.Println("Result count: ", len(as))
}

func TestSqlite_DBX_Insert_1(t *testing.T) {
	db := getSqliteDBX()

	r1 := newRandomRecordForApp("dbx")

	err := db.Model(r1).Insert()
	require.NoError(t, err)
}
