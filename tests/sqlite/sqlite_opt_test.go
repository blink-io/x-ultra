package sqlite

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"os"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/dbr"
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
	db, err := xsql.NewDBX(sqliteCfg(),
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
	db, err := xsql.NewDBR(sqliteCfg(),
		xsql.DBREventReceiver(dbr.NewTimingEventReceiver()),
	)

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
