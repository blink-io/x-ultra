package sql

import (
	"github.com/pocketbase/dbx"
)

const (
	AccessorDBX = "dbx"
)

type (
	idbx = dbx.DB

	DBX struct {
		*idbx
		accessor string
	}
)

func NewDBX(o *Options) (*DBX, error) {
	o = setupOptions(o)
	o.accessor = AccessorDBX

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	rdb := dbx.NewFromDB(sqlDB, o.Dialect)
	rdb.LogFunc = o.Logger
	db := &DBX{
		idbx:     rdb,
		accessor: o.accessor,
	}
	return db, nil
}

func (d *DBX) Accessor() string {
	return d.accessor
}
