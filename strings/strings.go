package strings

import (
	"strings"
)

// SplitQualifiedName Splits a fully qualified name and returns its namespace and name.
// Assumes that the input 'str' has been validated.
func SplitQualifiedName(str string) (string, string) {
	parts := strings.Split(str, "/")
	if len(parts) < 2 {
		return "", str
	}
	return parts[0], parts[1]
}

// Shorten returns the first N slice of a string.
func Shorten(str string, n int) string {
	if len(str) <= n {
		return str
	}
	return str[:n]
}
