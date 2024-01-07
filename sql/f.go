package sql

import (
	"github.com/blink-io/x/sql/f"
)

func F() f.F {
	return f.New()
}

func (db *DB) F() f.F {
	return F()
}
