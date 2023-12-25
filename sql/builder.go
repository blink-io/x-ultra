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

func (d *DBX) B() Builder {
	return B()
}

func (d *DBQ) B() Builder {
	return B()
}

func (d *DBW) B() Builder {
	return B()
}

func (d *DBR) B() Builder {
	return B()
}

func (d *DBP) B() Builder {
	return B()
}

func (d *DBM) B() Builder {
	return B()
}
