package db

import (
	"context"
	"database/sql"
	"io"

	xsql "github.com/blink-io/x/sql"

	"github.com/uptrace/bun"
)

const (
	Accessor = "db(bun)"
)

type (
	idb = bun.DB

	IDB interface {
		bun.IDB
		io.Closer
		ToSQL() string
	}

	Config = xsql.Config

	DB struct {
		*idb
		sqlDB    *sql.DB
		info     xsql.DBInfo
		accessor string
		rawName  string
	}
)

var _ xsql.HealthChecker = (*DB)(nil)
var _ IDB = (*DB)(nil)

func New(c *Config, ops ...Option) (*DB, error) {
	c = xsql.SetupConfig(c)

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
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: Accessor,
		info:     c.DBInfo(),
	}

	return db, nil
}

func (db *DB) RegisterModel(m ...any) {
	db.idb.RegisterModel(m...)
}

func (db *DB) ToSQL() string {
	return db.idb.String()
}

func (db *DB) Close() error {
	if db.idb != nil {
		return db.idb.Close()
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
