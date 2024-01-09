package sqlite

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/blink-io/x/bun/extra/timing"
	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/blink-io/x/sql/driver/hooks"
	logginghook "github.com/blink-io/x/sql/driver/hooks/logging"
)

var ctx = context.Background()

var sqlitePath = filepath.Join(".", "bun_demo.db")

func dbOpts() []xdb.Option {
	opts := []xdb.Option{
		xdb.WithQueryHooks(timing.New()),
	}
	return opts
}

func sqliteCfg() *xsql.Config {
	var cfg = &xsql.Config{
		Dialect:     xsql.DialectSQLite,
		Host:        sqlitePath,
		DriverHooks: newDriverHooks(),
		Logger: func(format string, args ...any) {
			msg := fmt.Sprintf(format, args...)
			slog.Default().With("db", "sqlite").Info(msg, "mode", "test")
		},
		Loc: time.Local,
	}
	return cfg
}

func newDriverHooks() []hooks.Hooks {
	hs := []hooks.Hooks{
		logginghook.Func(func(format string, args ...any) {
			slog.Default().Info(fmt.Sprintf(format, args...))
		}),
	}
	return hs
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
