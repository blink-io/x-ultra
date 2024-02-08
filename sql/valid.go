package sql

import (
	"database/sql"
	"time"
)

func ValidNull[T any](v T) sql.Null[T] {
	return sql.Null[T]{V: v, Valid: true}
}

// ValidBool is short for valid NullBool
func ValidBool(v bool) sql.NullBool {
	return sql.NullBool{Bool: v, Valid: true}
}

// ValidByte is short for valid NullByte
func ValidByte(v byte) sql.NullByte {
	return sql.NullByte{Byte: v, Valid: true}
}

// ValidFloat64 is short for valid NullFloat64
func ValidFloat64(v float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: v, Valid: true}
}

// ValidInt16 is short for valid NullFloat64
func ValidInt16(v int16) sql.NullInt16 {
	return sql.NullInt16{Int16: v, Valid: true}
}

// ValidInt32 is short for valid NullInt32
func ValidInt32(v int32) sql.NullInt32 {
	return sql.NullInt32{Int32: v, Valid: true}
}

// ValidInt64 is short for valid NullInt64
func ValidInt64(v int64) sql.NullInt64 {
	return sql.NullInt64{Int64: v, Valid: true}
}

// ValidString is short for valid NullString
func ValidString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: true}
}

// ValidTime is short for valid NullTime
func ValidTime(v time.Time) sql.NullTime {
	return sql.NullTime{Time: v, Valid: true}
}
