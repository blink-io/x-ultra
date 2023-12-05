package attrs

import (
	"log/slog"
)

type Option func(*options)

type options struct {
	fields []any
}

func applyOptions(ops ...Option) *options {
	opt := &options{
		fields: make([]any, 0),
	}
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func Append(key string, value any) Option {
	return func(o *options) {
		o.fields = append(o.fields, slog.Any(key, value))
	}
}
