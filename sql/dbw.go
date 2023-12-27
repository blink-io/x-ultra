package sql

import (
	"github.com/ilibs/gosql/v2"
)

const (
	AccessorDBW = "dbw"
)

type (
	idbw = gosql.DB

	DBW struct {
		*idbw
		accessor string
	}
)

func NewDBW(o *Options) (*DBW, error) {
	o = setupOptions(o)
	o.accessor = AccessorDBW

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	rdb := gosql.OpenWithDB(o.Dialect, sqlDB)

	db := &DBW{
		idbw:     rdb,
		accessor: o.accessor,
	}
	return db, nil
}

func (d *DBW) Accessor() string {
	return d.accessor
}
