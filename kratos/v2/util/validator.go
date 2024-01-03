package util

import (
	"context"
)

type Validator interface {
	Validate(context.Context) error
}
