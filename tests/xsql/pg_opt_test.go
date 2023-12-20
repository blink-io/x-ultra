package bun

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	"github.com/doug-martin/goqu/v9"
)

func getPgSqlDB() *sql.DB {
	db, err1 := xsql.NewSqlDB(pgOpt)

	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDB() *xsql.DB {
	db, err1 := xsql.NewDB(pgOpt)
	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDBX() *xsql.DBX {
	db, err1 := xsql.NewDBX(pgOpt)
	//db.AddQueryHook(logging.Func(log.Printf))
	if err1 != nil {
		panic(err1)
	}

	handleDBX(db)

	return db
}

func getPgGoquDB() *goqu.Database {
	rdb := getPgSqlDB()
	db := goqu.New(xsql.DialectPostgres, rdb)
	handleDBQ(db)

	return db
}

func getPgDBR() *xsql.DBR {
	db, err := xsql.NewDBR(pgOpt)

	if err != nil {
		panic(err)
	}

	return db
}
