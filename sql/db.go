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
		accessor string
	}
)

func NewDB(o *Options) (*DB, error) {
	o = setupOptions(o)
	o.accessor = AccessorDB

	sd, sqlDB, err := GetDialect(o)
	if err != nil {
		return nil, err
	}

	rdb := bun.NewDB(sqlDB, sd, bun.WithDiscardUnknownColumns())
	if queryHooks := o.QueryHooks; len(queryHooks) > 0 {
		for _, h := range queryHooks {
			rdb.AddQueryHook(h)
		}
	}

	db := &DB{
		idb:      rdb,
		accessor: o.accessor,
	}

	return db, nil
}

func (d *DB) Accessor() string {
	return d.accessor
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
