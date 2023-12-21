package sql

import (
	"database/sql"
	"time"
)

// InvalidBool is short for Invalid Bool
func InvalidBool(v bool) sql.NullBool {
	return sql.NullBool{Bool: v, Valid: false}
}

// InInvalidByte is short for Invalid InvalidByte
func InInvalidByte(v byte) sql.NullByte {
	return sql.NullByte{Byte: v, Valid: false}
}

// InInvalidFloat64 is short for Invalid InvalidFloat64
func InInvalidFloat64(v float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: v, Valid: false}
}

// InvalidInt16 is short for Invalid InvalidFloat64
func InvalidInt16(v int16) sql.NullInt16 {
	return sql.NullInt16{Int16: v, Valid: false}
}

// InvalidInt32 is short for Invalid InvalidInt32
func InvalidInt32(v int32) sql.NullInt32 {
	return sql.NullInt32{Int32: v, Valid: false}
}

// InvalidInt64 is short for Invalid InvalidInt64
func InvalidInt64(v int64) sql.NullInt64 {
	return sql.NullInt64{Int64: v, Valid: false}
}

// InvalidString is short for Invalid InvalidString
func InvalidString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: false}
}

// InvalidTime is short for Invalid InvalidTime
func InvalidTime(v time.Time) sql.NullTime {
	return sql.NullTime{Time: v, Valid: false}
}
