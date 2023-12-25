package sql

import (
	"github.com/jmoiron/sqlx"
)

const (
	AccessorDBW = "dbw"
)

type (
	idbw = sqlx.DB

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

	rdb := sqlx.NewDb(sqlDB, o.Dialect)

	db := &DBW{
		idbw:     rdb,
		accessor: o.accessor,
	}
	return db, nil
}

func (d *DBW) Accessor() string {
	return d.accessor
}
