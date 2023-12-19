package bun

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"testing"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/hooks"
	"github.com/stretchr/testify/require"
)

func getDBX(t *testing.T) *xsql.DBX {
	dbPath := filepath.Join(".", "bun_demo.db")

	fmt.Println("db path: ", dbPath)

	db, err1 := xsql.NewDBX(&xsql.Options{
		Dialect:     xsql.DialectSQLite,
		Host:        dbPath,
		DriverHooks: []hooks.Hooks{
			//timing.New(timing.Logf(func(format string, args ...any) {
			//	msg := fmt.Sprintf(format, args...)
			//	slog.Default().Info(msg)
			//})),
		},
	})
	//db.AddQueryHook(logging.Func(log.Printf))
	require.NoError(t, err1)

	return db
}

func TestSQLite3_Select_Funcs_DBX(t *testing.T) {
	db := getDBX(t)
	db.LogFunc = func(format string, args ...interface{}) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}

	sqlF := "select %s as payload"
	funcs := []string{
		"hex(randomblob(32))",
		"random()",
		"sqlite_version()",
		"total_changes()",
		`lower("HELLO")`,
		`upper("hello")`,
		`length("hello")`,
		`length("我是世界")`,
		//`concat("Hello", ",", "World")`,
	}

	for _, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		q := db.NewQuery(ss)
		var v string
		require.NoError(t, q.Row(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}
