package dbk

import (
	"context"
	"database/sql"
	"io"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/dbk/adapters/kstd"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/sqldialect"
)

const (
	Accessor = "dbk(ksql)"
)

type (
	idb = ksql.DB

	IDB interface {
		DBF

		io.Closer

		xsql.IDBExt
	}

	DB struct {
		idb
		sqlDB *sql.DB
		info  xsql.DBInfo
	}
)

var _ IDB = (*DB)(nil)

func New(c *xsql.Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)
	c.Accessor = Accessor
	dialect := xsql.GetFormalDialect(c.Dialect)

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	var dp sqldialect.Provider
	switch dialect {
	case xsql.DialectPostgres:
		dp = sqldialect.PostgresDialect{}
	case xsql.DialectMySQL:
		dp = sqldialect.MysqlDialect{}
	case xsql.DialectSQLite:
		dp = sqldialect.Sqlite3Dialect{}
	default:
		return nil, xsql.ErrUnsupportedDialect
	}

	opts := applyOptions(ops...)
	if opts != nil {

	}

	rdb, err := ksql.NewWithAdapter(kstd.NewSQLAdapter(sqlDB), dp)

	s := &DB{
		idb:   rdb,
		sqlDB: sqlDB,
		info:  xsql.NewDBInfo(c),
	}

	return s, nil
}

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}

func (db *DB) DBInfo() xsql.DBInfo {
	return db.info
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
