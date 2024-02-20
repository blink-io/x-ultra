package pg

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/blink-io/x/sql/dbq"
	"github.com/blink-io/x/sql/dbr"
	"github.com/blink-io/x/sql/dbx"
	"github.com/stephenafamo/bob"
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

func getPgDBQ() *dbq.DB {
	db, err := dbq.New(pgCfg())
	if err != nil {
		panic(err)
	}

	handleDBQ(db)

	return db
}

func getPgDBX() *dbr.DB {
	db, err := dbr.New(pgCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getPgDBZ() *dbx.DB {
	ops := []dbx.Option{
		dbx.ExecWrappers(func(exec bob.Executor) bob.Executor {
			return dbx.ExecOnError(exec, func(e error) error {
				return xsql.WrapError(e)
			})
		}),
	}
	db, err := dbx.New(pgCfg(), ops...)

	if err != nil {
		panic(err)
	}

	return db
}
