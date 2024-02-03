package pg

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/blink-io/x/sql/dbm"
	"github.com/blink-io/x/sql/dbq"
	"github.com/blink-io/x/sql/dbr"
	"github.com/blink-io/x/sql/dbx"
	"github.com/blink-io/x/sql/dbz"
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

func getPgDBX() *dbx.DB {
	db, err1 := dbx.New(pgCfg(),
		dbx.WithExecLogFunc(func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
			slog.Default().Info("dbx exec log",
				slog.String("sql", sql),
				slog.Duration("time", t),
			)
		}),
		dbx.WithQueryLogFunc(func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
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

func getPgDBQ() *dbq.DB {
	db, err := dbq.New(pgCfg())
	if err != nil {
		panic(err)
	}

	handleDBQ(db)

	return db
}

func getPgDBR() *dbr.DB {
	db, err := dbr.New(pgCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getPgDBM() *dbm.DB {
	db, err := dbm.New(pgCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getPgDBZ() *dbz.DB {
	ops := []dbz.Option{
		dbz.ExecWrappers(func(exec bob.Executor) bob.Executor {
			return dbz.ExecOnError(exec, func(e error) error {
				return xsql.WrapError(e)
			})
		}),
	}
	db, err := dbz.New(pgCfg(), ops...)

	if err != nil {
		panic(err)
	}

	return db
}
