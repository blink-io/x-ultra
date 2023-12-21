package xsql

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
)

func getPgSqlDB() *sql.DB {
	db, err1 := xsql.NewSqlDB(pgOpt())

	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDB() *xsql.DB {
	db, err1 := xsql.NewDB(pgOpt())
	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDBX() *xsql.DBX {
	db, err1 := xsql.NewDBX(pgOpt())
	//db.AddQueryHook(logging.Func(log.Printf))
	if err1 != nil {
		panic(err1)
	}

	handleDBX(db)

	return db
}

func getPgDBQ() *xsql.DBQ {
	db, err := xsql.NewDBQ(pgOpt())
	if err != nil {
		panic(err)
	}

	handleDBQ(db)

	return db
}

func getPgDBR() *xsql.DBR {
	db, err := xsql.NewDBR(pgOpt())

	if err != nil {
		panic(err)
	}

	return db
}

func getPgDBM() *xsql.DBM {
	db, err := xsql.NewDBM(pgOpt())

	if err != nil {
		panic(err)
	}

	return db
}
