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

func (db *DB) B() Builder {
	return B()
}

func (db *DBX) B() Builder {
	return B()
}

func (db *DBQ) B() Builder {
	return B()
}

func (db *DBW) B() Builder {
	return B()
}

func (db *DBR) B() Builder {
	return B()
}

func (db *DBP) B() Builder {
	return B()
}

func (db *DBM) B() Builder {
	return B()
}
