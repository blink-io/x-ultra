package bun

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/hooks"
	"github.com/blink-io/x/sql/hooks/timing"
	"github.com/doug-martin/goqu/v9"
)

var Debug = false

var ctx = context.Background()

var pgOpt = &xsql.Options{
	Context:       ctx,
	Dialect:       xsql.DialectPostgres,
	Name:          "blink",
	User:          "blinkbot",
	Port:          15432,
	Host:          "192.168.11.179",
	Loc:           time.Local,
	ValidationSQL: "SELECT 1;",
	ClientName:    "blink-dev",
	DriverHooks:   newDriverHooks(),
}

func init() {
	pwd := getPwd()
	pgOpt.Password = pwd
}

func getPwd() string {
	homedir, _ := os.UserHomeDir()

	data, err := os.ReadFile(filepath.Join(homedir, ".passwd.pg"))
	if err != nil {
		panic(err)
	}

	// Remove \n
	pwd := string(data[:len(data)-1])
	return pwd
}

func getDBWithPG() *xsql.DB {
	db, err1 := xsql.NewDB(pgOpt)
	if err1 != nil {
		panic(err1)
	}

	return db
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

func getSqlDBWithSQLite() *sql.DB {
	dbPath := filepath.Join(".", "bun_demo.db")

	fmt.Println("db path: ", dbPath)

	db, err1 := xsql.NewSqlDB(&xsql.Options{
		Dialect:     xsql.DialectSQLite,
		Host:        dbPath,
		DriverHooks: newDriverHooks(),
	})
	//db.AddQueryHook(logging.Func(log.Printf))
	if err1 != nil {
		panic(err1)
	}

	return db
}

func getDBWithSQLite() *xsql.DB {
	dbPath := filepath.Join(".", "bun_demo.db")

	fmt.Println("db path: ", dbPath)

	db, err1 := xsql.NewDB(&xsql.Options{
		Dialect:     xsql.DialectSQLite,
		Host:        dbPath,
		DriverHooks: newDriverHooks(),
	})
	//db.AddQueryHook(logging.Func(log.Printf))
	//db.AddQueryHook(timing.New())

	if err1 != nil {
		panic(err1)
	}

	return db
}

func getDBXWithSQLite() *xsql.DBX {
	dbPath := filepath.Join(".", "bun_demo.db")

	fmt.Println("db path: ", dbPath)

	db, err1 := xsql.NewDBX(&xsql.Options{
		Dialect:     xsql.DialectSQLite,
		Host:        dbPath,
		DriverHooks: newDriverHooks(),
	})
	//db.AddQueryHook(logging.Func(log.Printf))
	if err1 != nil {
		panic(err1)
	}

	handleDBX(db)

	return db
}

func getGoquDBWithSQLite() *goqu.Database {
	rdb := getSqlDBWithSQLite()

	db := goqu.New(xsql.DialectSQLite, rdb)
	handleGoqu(db)

	return db
}

type logger func(format string, args ...any)

func (l logger) Printf(format string, args ...any) {
	l(format, args...)
}

var _ goqu.Logger = (logger)(nil)

func handleGoqu(db *goqu.Database) {
	db.Logger(logger(func(format string, args ...any) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}))
}

func handleDBX(db *xsql.DBX) {
	db.LogFunc = func(format string, args ...interface{}) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}
}
