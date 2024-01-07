package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

const (
	AccessorDB = "db(bun)"
	RawNameDB  = "bun"

	DefaultTimeout = 15 * time.Second
)

type (
	idb = bun.DB

	IDB = bun.IDB

	DB struct {
		*idb
		sqlDB    *sql.DB
		info     DBInfo
		accessor string
		rawName  string
	}
)

var _ HealthChecker = (*DB)(nil)

func NewDB(c *Config, ops ...DBOption) (*DB, error) {
	c = setupConfig(c)
	c.accessor = AccessorDB

	dialectOpts := make([]DialectOption, 0)
	if c.Loc != nil {
		dialectOpts = append(dialectOpts, DialectWithLoc(c.Loc))
	}
	sd, sqlDB, err := GetDialect(c, dialectOpts...)
	if err != nil {
		return nil, err
	}

	rdb := bun.NewDB(sqlDB, sd, bun.WithDiscardUnknownColumns())

	dbOpts := applyDBOptions(ops...)
	for _, h := range dbOpts.queryHooks {
		rdb.AddQueryHook(h)
	}

	db := &DB{
		idb:      rdb,
		sqlDB:    sqlDB,
		accessor: c.accessor,
		rawName:  RawNameDB,
		info:     newDBInfo(c),
	}

	return db, nil
}

func (db *DB) Accessor() string {
	return db.accessor
}

func (db *DB) DBInfo() DBInfo {
	return db.info
}

func (db *DB) RegisterModel(m any) {
	db.idb.RegisterModel(m)
}

func (db *DB) Close() error {
	if db.idb != nil {
		return db.idb.Close()
	}
	return nil
}

func (db *DB) HealthCheck(ctx context.Context) error {
	return doPingFunc(ctx, db.sqlDB.PingContext)
}
