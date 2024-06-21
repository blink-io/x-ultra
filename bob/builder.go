package bob

import (
	xsql "github.com/blink-io/x/sql"
)

type Builder = xsql.Builder

func (db *DB) B() Builder {
	return xsql.B()
}
