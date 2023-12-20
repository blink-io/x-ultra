package sql

import (
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
)

type (
	dbrc = dbr.Connection
)

type DBR struct {
	*dbrc
}

func NewDBR(o *Options) (*DBR, error) {
	sqlDB, err := NewSqlDB(o)
	if err != nil {
		return nil, err
	}
	var d dbr.Dialect
	switch o.Dialect {
	case "mysql":
		d = dialect.MySQL
	case "postgres", "pgx":
		d = dialect.PostgreSQL
	case "sqlite3", "sqlite":
		d = dialect.SQLite3
	case "mssql":
		d = dialect.MSSQL
	default:
		return nil, dbr.ErrNotSupported
	}
	cc := &dbrc{
		DB:      sqlDB,
		Dialect: d,
	}
	db := &DBR{
		dbrc: cc,
	}
	return db, nil
}
