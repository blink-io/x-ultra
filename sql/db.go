package sql

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type (
	db = bun.DB

	DB struct {
		*db
	}

	DBer interface {
		DB() *DB
	}
)

func NewDB(o *Options) (*DB, error) {
	o, err := setupOptions(o)
	if err != nil {
		return nil, err
	}

	sd, sqlDB, err := GetDialect(o)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, sd, bun.WithDiscardUnknownColumns())
	if queryHooks := o.QueryHooks; len(queryHooks) > 0 {
		for _, h := range queryHooks {
			db.AddQueryHook(h)
		}
	}

	return &DB{db: db}, nil
}

func (d *DB) RegisterModel(m any) {
	d.db.RegisterModel(m)
}

func (d *DB) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

func (d *DB) Raw() *sql.DB {
	return d.db.DB
}
