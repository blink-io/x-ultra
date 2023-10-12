package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var _ validator.Func = isUsername

var (
	usernameRegexString = "^[a-zA-Z]{1}[a-zA-Z0-9_]+$"
	usernameRegex       = regexp.MustCompile(usernameRegexString)
)

func isUsername(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}
