package sql

import (
	"database/sql"
	"errors"
)

type Checker interface {
	IsMySQL() bool
	IsPostgres() bool
	IsSQLite() bool
}

// IsNoRows .
func IsNoRows(e error) bool {
	return errors.Is(e, sql.ErrNoRows)
}
