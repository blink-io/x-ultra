//go:build misc

package sql

import (
	sq "github.com/Masterminds/squirrel"
)

type Builder = sq.StatementBuilderType

var sb = sq.StatementBuilder

// B is short for SQL Builder
func B() Builder {
	return sb
}

func (d *DB) B() Builder {
	return B()
}
