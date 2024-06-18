package model

type Schema[T any, C any] struct {
	Type        *T
	PK          string
	Label       string
	Alias       string
	TableName   string
	ColumnNames C
}
