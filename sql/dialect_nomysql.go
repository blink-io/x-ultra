//go:build !mysql

package sql

func IsMySQLConstraintCodes(n uint16) bool {
	return false
}
