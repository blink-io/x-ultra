package sql

import (
	"github.com/uptrace/bun"
)

type (
	idb = bun.DB

	IDB = bun.IDB

	DB struct {
		*idb
	}
)

func NewDB(o *Options) (*DB, error) {
	o = setupOptions(o)
	o.accessor = "bun " + bun.Version()

	sd, sqlDB, err := GetDialect(o)
	if err != nil {
		return nil, err
	}

	idb := bun.NewDB(sqlDB, sd, bun.WithDiscardUnknownColumns())
	if queryHooks := o.QueryHooks; len(queryHooks) > 0 {
		for _, h := range queryHooks {
			idb.AddQueryHook(h)
		}
	}

	return &DB{idb: idb}, nil
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
