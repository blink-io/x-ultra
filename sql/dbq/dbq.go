package dbq

import (
	"context"
	"database/sql"
	"io"
	"log/slog"

	xsql "github.com/blink-io/x/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

const (
	Accessor = "dbq(goqu)"
)

func init() {
	// Overrides default dialects
	goqu.RegisterDialect(xsql.DialectPostgres, postgres.DialectOptions())
	goqu.RegisterDialect(xsql.DialectMySQL, mysql.DialectOptions())
	//
	sqliteDialectOpts := sqlite3.DialectOptions()
	// See: https://www.sqlite.org/lang_returning.html
	// The RETURNING syntax has been supported by SQLite since version 3.35.0 (2021-03-12).
	sqliteDialectOpts.SupportsReturn = true
	goqu.RegisterDialect(xsql.DialectSQLite, sqliteDialectOpts)
}

type (
	idb = goqu.Database

	Config = xsql.Config

	// IDB defines
	IDB interface {
		DBF

		io.Closer

		xsql.IDBExt
	}

	DB struct {
		*idb
		sqlDB    *sql.DB
		accessor string
		info     xsql.DBInfo
	}
)

var _ IDB = (*DB)(nil)

func New(c *xsql.Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)
	c.Accessor = Accessor

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	if c.Loc != nil {
		goqu.SetTimeLocation(c.Loc)
		slog.Info("Setup global time location for goqu", slog.String("location", c.Loc.String()))
	}

	opts := applyOptions(ops...)
	if opts != nil {
	}

	rdb := goqu.New(c.Dialect, sqlDB)
	if c.Logger != nil {
		rdb.Logger(xsql.PrintfLogger(c.Logger))
	}
	db := &DB{
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
		info:     c.DBInfo(),
	}
	return db, nil
}

func (db *DB) Close() error {
	return db.sqlDB.Close()
}

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}

func (db *DB) DBInfo() xsql.DBInfo {
	return db.info
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
