package sqlite

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBW_Select_Funcs(t *testing.T) {
	db := getSqliteDBW()

	sqlF := "select %s as payload where ? = ?"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		var rt string
		err1 := db.Get(&rt, ss, 1, 1)
		require.NoError(t, err1)
		fmt.Println(k, "=>", rt)
	}
}

func TestSqlite_DBW_Insert_1(t *testing.T) {
	db := getSqliteDBW()

	sql := "insert into applications (id, iid, name,status, code, type, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8)"

	args := []any{
		id.ShortUUID(),
		gofakeit.Int32(),
		"from-" + db.Accessor() + "-" + gofakeit.Name(),
		"ok",
		"001-" + id.ShortID(),
		"type-01",
		time.Now(),
		time.Now(),
	}
	//r1 := newRandomRecordForApp("dbp")

	_, err := db.Exec(sql, args...)
	require.NoError(t, err)
}
