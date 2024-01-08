package pg

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

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
	db, err1 := xsql.NewDBX(pgCfg(),
		xsql.DBXExecLogFunc(func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
			slog.Default().Info("dbx exec log",
				slog.String("sql", sql),
				slog.Duration("time", t),
			)
		}),
		xsql.DBXQueryLogFunc(func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
			slog.Default().Info("dbx query log",
				slog.String("sql", sql),
				slog.Duration("time", t),
			)
		}),
	)
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
