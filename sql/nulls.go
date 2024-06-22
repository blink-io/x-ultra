package sql

import (
	"database/sql"
)

func ValidNull[T any](v T) sql.Null[T] {
	return sql.Null[T]{V: v, Valid: true}
}

func InvalidNull[T any](v T) sql.Null[T] {
	return sql.Null[T]{V: v, Valid: false}
}
