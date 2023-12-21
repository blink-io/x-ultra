package xsql_test

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	xsql "github.com/blink-io/x/sql"
)

var sqlitePath = filepath.Join(".", "bun_demo.db")

var sqliteOpts = &xsql.Options{
	Dialect: xsql.DialectSQLite,
	Host:    sqlitePath,
	DOptions: []xsql.DOption{
		xsql.WithLocation(time.Local),
	},
	DriverHooks: newDriverHooks(),
	Logger: func(format string, args ...any) {
		slog.Info(fmt.Sprintf(format, args...))
	},
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