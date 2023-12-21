package sql

import (
	"github.com/pocketbase/dbx"
)

type (
	idbx = dbx.DB
	DBX  struct {
		*idbx
	}
)

func NewDBX(o *Options) (*DBX, error) {
	o = setupOptions(o)
	o.accessor = "dbx"

	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}

	rdb := dbx.NewFromDB(sqlDB, o.Dialect)
	rdb.LogFunc = o.Logger
	db := &DBX{
		idbx: rdb,
	}
	return db, nil
}
