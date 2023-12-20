package bun

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/doug-martin/goqu/v9/sqlgen"
)

func getSqliteSqlDB() *sql.DB {
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

func getSqliteDB() *xsql.DB {
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

func getSqliteDBX() *xsql.DBX {
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

func getSqliteGoquDB() *goqu.Database {
	goqu.RegisterDialect(xsql.DialectSQLite, sqlite3.DialectOptions())
	sqlgen.SetTimeLocation(time.Local)

	rdb := getSqliteSqlDB()
	db := goqu.New(xsql.DialectSQLite, rdb)
	handleGoqu(db)

	return db
}

func getSqliteDBR() *xsql.DBR {
	db, err1 := xsql.NewDBR(sqliteOpts)
	//db.AddQueryHook(logging.Func(log.Printf))
	if err1 != nil {
		panic(err1)
	}

	return db
}
