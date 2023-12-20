package bun

import (
	"database/sql"
	"time"

	xsql "github.com/blink-io/x/sql"

	"github.com/doug-martin/goqu/v9/sqlgen"
)

func getSqliteSqlDB() *sql.DB {
	db, err := xsql.NewSqlDB(sqliteOpts)
	//db.AddQueryHook(logging.Func(log.Printf))
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDB() *xsql.DB {
	db, err := xsql.NewDB(sqliteOpts)
	//db.AddQueryHook(logging.Func(log.Printf))
	//db.AddQueryHook(timing.New())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBX() *xsql.DBX {
	db, err := xsql.NewDBX(sqliteOpts)
	//db.AddQueryHook(logging.Func(log.Printf))
	if err != nil {
		panic(err)
	}

	handleDBX(db)

	return db
}

func getSqliteGoquDB() *xsql.DBQ {
	sqlgen.SetTimeLocation(time.Local)

	db, err := xsql.NewDBQ(sqliteOpts)
	if err != nil {
		panic(err)
	}

	handleDBQ(db)

	return db
}

func getSqliteDBR() *xsql.DBR {
	db, err := xsql.NewDBR(sqliteOpts)

	if err != nil {
		panic(err)
	}

	return db
}
