package validator

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"go.uber.org/multierr"
)

var vd = validator.New()

func init() {
	err1 := vd.RegisterValidation("username", isUsername)

	if err := multierr.Combine(err1); err != nil {
		slog.Warn("validator_v10: register validations failed", slog.String("error", err.Error()))
	}
}

func Default() *validator.Validate {
	return vd
}

var _ validator.Func = isUsername

func isUsername(fl validator.FieldLevel) bool {
	return IsUsername(fl.Field().String())
}
