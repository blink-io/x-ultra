package bun

import (
	"context"
	"fmt"
	"log/slog"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/hooks"
	"github.com/blink-io/x/sql/hooks/timing"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

var Debug = false

func DebugEnabled() {
	Debug = true
}

var ctx = context.Background()

func init() {
	doGoquInit()
}

func doGoquInit() {
	goqu.RegisterDialect(xsql.DialectSQLite, sqlite3.DialectOptions())
	goqu.RegisterDialect(xsql.DialectPostgres, postgres.DialectOptions())
	goqu.RegisterDialect(xsql.DialectMySQL, mysql.DialectOptions())
	fmt.Println("Invoke goqu init")
}

func newDriverHooks() []hooks.Hooks {
	drvHooks := make([]hooks.Hooks, 0)

	if Debug {
		drvHooks = append(drvHooks, timing.New(timing.Logf(func(format string, args ...any) {
			msg := fmt.Sprintf(format, args...)
			slog.Default().Info(msg)
		})))
	}
	return drvHooks
}
