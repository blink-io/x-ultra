package model

type Meta[T any, C any] struct {
	PK      string
	Label   string
	Table   string
	Alias   string
	Type    *T
	Columns *C
}
