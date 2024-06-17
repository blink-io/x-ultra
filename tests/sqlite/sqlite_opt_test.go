package sqlite

import (
	"database/sql"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
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
