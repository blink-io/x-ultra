package model

type Schema[M any, C any] struct {
	PK      string
	Label   string
	Alias   string
	Table   string
	Model   *M
	Columns C
}

type Column string
