package sqlitedialect

type options struct {
	utc bool
}

type Option func(o *options)

func applyOptions(ops ...Option) *options {
	opt := new(options)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

func UTC() Option {
	return func(o *options) {
		o.utc = true
	}
}
