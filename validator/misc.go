package validator

import (
	"regexp"
)

const Username = "^[a-zA-Z]{1}[a-zA-Z0-9_]+$"

var (
	usernameRegex = regexp.MustCompile(Username)
)

func init() {
	TagMap["username"] = IsUsername
}

// IsUsername returns true if value is username
func IsUsername(value string) bool {
	return usernameRegex.MatchString(value)
}
