package dbr

import (
	"context"
	"database/sql"
	"io"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/dbr/dialect"
	"github.com/gocraft/dbr/v2"
)

const (
	Accessor = "dbr(dbr)"

	DefaultTimeout = 15 * time.Second
)

type (
	idb = dbr.Session

	Config = xsql.Config

	IDB interface {
		DBF

		io.Closer

		xsql.IDBExt
	}

	DB struct {
		*idb
		sqlDB    *sql.DB
		accessor string
		rawName  string
		info     xsql.DBInfo
	}
)

var _ IDB = (*DB)(nil)

func New(c *Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)
	c.Accessor = Accessor

	sqlDB, err := xsql.NewSqlDB(c)
	if err != nil {
		return nil, err
	}

	var d dbr.Dialect
	switch c.Dialect {
	case xsql.DialectMySQL:
		d = dialect.MySQL
	case xsql.DialectPostgres:
		d = dialect.Postgres
	case xsql.DialectSQLite:
		d = dialect.SQLite3
	default:
		return nil, xsql.ErrUnsupportedDialect
	}

	opts := applyOptions(ops...)

	var er dbr.EventReceiver
	if er = opts.er; er == nil {
		er = new(dbr.NullEventReceiver)
	}

	cc := &dbr.Connection{
		DB:            sqlDB,
		Dialect:       d,
		EventReceiver: er,
	}
	rdb := cc.NewSession(er)
	rdb.Timeout = DefaultTimeout

	db := &DB{
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
		info:     c.DBInfo(),
	}
	return db, nil
}

func (db *DB) DBInfo() xsql.DBInfo {
	return db.info
}

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
