package attrs

type Option func(*options)

type options struct {
	fields []any
}

func applyOptions(ops ...Option) *options {
	opt := new(options)
	for _, o := range ops {
		o(opt)
	}
	return opt
}
