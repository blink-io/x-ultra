package strings

import (
	"strings"
)

// StrictTrue defines strict true with these values: 1 yes true on
func StrictTrue(s string) bool {
	// true values: 1 yes true on
	switch ls := strings.ToLower(s); ls {
	case "1", "yes", "true", "on":
		return true
	default:
		return false
	}
}

// StrictFalse defines strict false with these values: 0 no false off
func StrictFalse(s string) bool {
	// false values: 0 no false off
	switch ls := strings.ToLower(s); ls {
	case "0", "no", "false", "off":
		return true
	default:
		return false
	}
}
