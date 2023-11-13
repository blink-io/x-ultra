package validator

import (
	"github.com/go-playground/validator/v10"
)

var _ validator.Func = isUsername

func isUsername(fl validator.FieldLevel) bool {
	return IsUsername(fl.Field().String())
}
