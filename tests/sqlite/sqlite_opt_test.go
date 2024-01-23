package sqlite

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"os"
	"time"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/blink-io/x/sql/dbk"
	"github.com/blink-io/x/sql/dbm"
	"github.com/blink-io/x/sql/dbp"
	"github.com/blink-io/x/sql/dbq"
	"github.com/blink-io/x/sql/dbr"
	"github.com/blink-io/x/sql/dbs"
	"github.com/blink-io/x/sql/dbw"
	"github.com/blink-io/x/sql/dbx"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

func init() {
	sqliteDialectOpts := sqlite3.DialectOptions()
	sqliteDialectOpts.SupportsReturn = true
	goqu.RegisterDialect(xsql.DialectSQLite, sqliteDialectOpts)
}

func getSqliteSqlDB() *sql.DB {
	db, err := xsql.NewSqlDB(sqliteCfg())
	//db.AddQueryHook(logging.Func(log.Printf))
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDB() *xdb.DB {
	db, err := xdb.New(sqliteCfg(), dbOpts()...)

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBX() *dbx.DB {
	db, err := dbx.New(sqliteCfg(),
		dbx.WithDbTag("db"),
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

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBQ() *dbq.DB {
	db, err := dbq.New(sqliteCfg())
	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBR() *dbr.DB {
	db, err := dbr.New(sqliteCfg(),
		dbr.WithEventReceiver(dbr.NewTimingEventReceiver()),
	)

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBM() *dbm.DB {
	db, err := dbm.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBP() *dbp.DB {
	db, err := dbp.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	db.TraceOn("[gorp]", log.New(os.Stdout, "dbp:", log.Lmicroseconds))

	return db
}

func getSqliteDBW() *dbw.DB {
	db, err := dbw.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBS() *dbs.DB {
	db, err := dbs.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}

func getSqliteDBK() *dbk.DB {
	db, err := dbk.New(sqliteCfg())

	if err != nil {
		panic(err)
	}

	return db
}
