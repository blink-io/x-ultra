package validator

import (
	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func init() {
	_ = RegisterValidation("username", isUsername)
}

func Default() *validator.Validate {
	return v
}

func RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func RegisterStructValidation(fn validator.StructLevelFunc, types ...any) {
	v.RegisterStructValidation(fn, types...)
}

func RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...any) {
	v.RegisterCustomTypeFunc(fn, types...)
}
