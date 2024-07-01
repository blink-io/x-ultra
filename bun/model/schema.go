package model

import rdb "github.com/blink-io/x/bun"

type Schema[M any, C any] struct {
	PK      string
	Label   string
	Alias   string
	Table   string
	Model   *M
	Columns C
}

type Column string

func (c Column) Name() rdb.Name {
	return rdb.Name(c)
}

func (c Column) Ident() rdb.Ident {
	return rdb.Ident(c)
}

func (c Column) Safe() rdb.Safe {
	return rdb.Safe(c)
}
