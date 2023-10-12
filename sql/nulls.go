package sql

import (
	"database/sql"
	"time"
)

// NullBool is short for valid NullBool
func NullBool(v bool) sql.NullBool {
	return sql.NullBool{Bool: v, Valid: true}
}

// NullByte is short for valid NullByte
func NullByte(v byte) sql.NullByte {
	return sql.NullByte{Byte: v, Valid: true}
}

// NullFloat64 is short for valid NullFloat64
func NullFloat64(v float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: v, Valid: true}
}

// NullInt16 is short for valid NullFloat64
func NullInt16(v int16) sql.NullInt16 {
	return sql.NullInt16{Int16: v, Valid: true}
}

// NullInt32 is short for valid NullInt32
func NullInt32(v int32) sql.NullInt32 {
	return sql.NullInt32{Int32: v, Valid: true}
}

// NullInt64 is short for valid NullInt64
func NullInt64(v int64) sql.NullInt64 {
	return sql.NullInt64{Int64: v, Valid: true}
}

// NullString is short for valid NullString
func NullString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: true}
}

// NullTime is short for valid NullTime
func NullTime(v time.Time) sql.NullTime {
	return sql.NullTime{Time: v, Valid: true}
}
