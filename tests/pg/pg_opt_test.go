package pg

import (
	"database/sql"

	xdb "github.com/blink-io/x/bun"

	xsql "github.com/blink-io/x/sql"
)

func getPgSqlDB() *sql.DB {
	db, err1 := xsql.NewSqlDB(pgCfg())

	if err1 != nil {
		panic(err1)
	}

	return db
}

func getPgDB() *xdb.DB {
	db, err1 := xdb.NewDB(pgCfg(), dbOpts()...)
	if err1 != nil {
		panic(err1)
	}

	return db
}
