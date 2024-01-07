package sqlite

import (
	"database/sql"
	"log"
	"os"

	xsql "github.com/blink-io/x/sql"
)

func getSqliteSqlDB() *sql.DB {
	db, err := xsql.NewSqlDB(sqliteCfg())
	//db.AddQueryHook(logging.Func(log.Printf))
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDB() *xsql.DB {
	db, err := xsql.NewDB(sqliteCfg(), dbOpts()...)

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBX() *xsql.DBX {
	db, err := xsql.NewDBX(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBQ() *xsql.DBQ {
	db, err := xsql.NewDBQ(sqliteCfg())
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBR() *xsql.DBR {
	db, err := xsql.NewDBR(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBM() *xsql.DBM {
	db, err := xsql.NewDBM(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBP() *xsql.DBP {
	db, err := xsql.NewDBP(sqliteCfg())

	if err != nil {
		panic(err)
	}

	db.TraceOn("[gorp]", log.New(os.Stdout, "dbp:", log.Lmicroseconds))

	return db
}

func getSqliteDBW() *xsql.DBW {
	db, err := xsql.NewDBW(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}
