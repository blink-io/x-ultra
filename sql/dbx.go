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

	sqlDB, err := GetSqlDB(o)
	if err != nil {
		return nil, err
	}

	idb := dbx.NewFromDB(sqlDB, o.Dialect)
	db := &DBX{
		idbx: idb,
	}
	return db, nil
}
