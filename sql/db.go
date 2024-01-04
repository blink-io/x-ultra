package sql

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	AccessorDB = "db(bun)"

	DefaultTimeout = 15 * time.Second
)

type (
	idb = bun.DB

	IDB = bun.IDB

	DB struct {
		*idb
		info     DBInfo
		accessor string
		name     string
	}

	DBInfo struct {
		Name    string
		Dialect string
	}
)

func newDBInfo(o *Options) DBInfo {
	return DBInfo{
		Name:    o.Name,
		Dialect: o.Dialect,
	}
}

func NewDB(o *Options) (*DB, error) {
	o = setupOptions(o)
	o.accessor = AccessorDB

	sd, sqlDB, err := GetDialect(o)
	if err != nil {
		return nil, err
	}

	rdb := bun.NewDB(sqlDB, sd, bun.WithDiscardUnknownColumns())
	for _, h := range o.QueryHooks {
		rdb.AddQueryHook(h)
	}

	db := &DB{
		idb:      rdb,
		accessor: o.accessor,
		info:     newDBInfo(o),
	}

	return db, nil
}

func (d *DB) Accessor() string {
	return d.accessor
}

func (d *DB) DBInfo() DBInfo {
	return DBInfo{
		Name: d.name,
	}
}

func (d *DB) RegisterModel(m any) {
	d.idb.RegisterModel(m)
}

func (d *DB) Close() error {
	if d.idb != nil {
		return d.idb.Close()
	}
	return nil
}
