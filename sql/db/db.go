package db

import (
	"context"
	"database/sql"
	"io"
	"reflect"

	xsql "github.com/blink-io/x/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

const (
	Accessor = "db(bun)"
)

type (
	Config = xsql.Config

	ext interface {
		RegisterModel(m ...any)

		Table(typ reflect.Type) *schema.Table
	}

	IDB interface {
		RawIDB

		io.Closer

		xsql.IDBExt

		ext
	}

	DB struct {
		*rdb
		sqlDB    *sql.DB
		info     xsql.DBInfo
		accessor string
		rawName  string
	}
)

var _ IDB = (*DB)(nil)

func New(c *Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)
	c.Accessor = Accessor

	dopts := make([]DialectOption, 0)
	if c.Loc != nil {
		dopts = append(dopts, DialectWithLoc(c.Loc))
	}
	sd, sqlDB, err := GetDialect(c, dopts...)
	if err != nil {
		return nil, err
	}

	rdb := bun.NewDB(sqlDB, sd, bun.WithDiscardUnknownColumns())

	opts := applyOptions(ops...)
	for _, h := range opts.queryHooks {
		rdb.AddQueryHook(h)
	}

	db := &DB{
		rdb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
		info:     c.DBInfo(),
	}

	return db, nil
}

func (db *DB) RegisterModel(m ...any) {
	db.rdb.RegisterModel(m...)
}

func (db *DB) SqlDB() *sql.DB {
	return db.sqlDB
}

func (db *DB) Close() error {
	if db.rdb != nil {
		return db.rdb.Close()
	}
	return nil
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) DBInfo() xsql.DBInfo {
	return db.info
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return xsql.DoPingContext(ctx, db.sqlDB)
}
