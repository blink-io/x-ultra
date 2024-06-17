package pg

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
)

func getPgSqlDB() *sql.DB {
	db, err1 := xsql.NewSqlDB(pgCfg())

	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDB() *xdb.DB {
	db, err1 := xdb.New(pgCfg(), dbOpts()...)
	if err1 != nil {
		panic(err1)
	}

	return db
}
