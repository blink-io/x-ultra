//go:build !sqlite3

package sql

func IsSQLiteConstraintCodes(code int) bool {
	return false
}
