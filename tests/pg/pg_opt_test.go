package pg

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
)

func getPgSqlDB() *sql.DB {
	db, err1 := xsql.NewSqlDB(pgCfg())

	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDB() *xsql.DB {
	db, err1 := xsql.NewDB(pgCfg(), dbOpts()...)
	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDBX() *xsql.DBX {
	db, err1 := xsql.NewDBX(pgCfg())
	//db.AddQueryHook(logging.Func(log.Printf))
	if err1 != nil {
		panic(err1)
	}

	handleDBX(db)

	return db
}

func getPgDBQ() *xsql.DBQ {
	db, err := xsql.NewDBQ(pgCfg())
	if err != nil {
		panic(err)
	}

	handleDBQ(db)

	return db
}

func getPgDBR() *xsql.DBR {
	db, err := xsql.NewDBR(pgCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getPgDBM() *xsql.DBM {
	db, err := xsql.NewDBM(pgCfg())

	if err != nil {
		panic(err)
	}

	return db
}
