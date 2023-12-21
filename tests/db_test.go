package tests

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"testing"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/stretchr/testify/require"
)

var sqlitePath = filepath.Join(".", "xsql", "bun_demo.db")

func sqliteOpts() *xsql.Options {
	var opt = &xsql.Options{
		Dialect: xsql.DialectSQLite,
		Host:    sqlitePath,
		DOptions: []xsql.DOption{
			xsql.WithLocation(time.Local),
		},
		//DriverHooks: newDriverHooks(),
		Logger: func(format string, args ...any) {
			//slog.SetDefault(custLogger)
			msg := fmt.Sprintf(format, args...)
			slog.Default().With("db", "sqlite").Info(msg, "mode", "test")
			//log, _ := zap.NewDevelopment()
			//if log != nil {
			//	log.Info(msg)
			//}
		},
	}
	return opt
}

func getSqliteDBX() *xsql.DBX {
	db, err := xsql.NewDBX(sqliteOpts())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteFuncMap() map[string]string {
	funcsMap := map[string]string{
		"hex":              "hex(randomblob(32))",
		"random":           "random()",
		"version":          "sqlite_version()",
		"changes":          "changes()",
		"total_changes":    "total_changes()",
		"lower":            `lower("HELLO")`,
		"upper":            `upper("hello")`,
		"length":           `length("hello")`,
		"sqlite_source_id": `sqlite_source_id()`,
		//`concat("Hello", ",", "World")`,
	}
	return funcsMap
}

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
