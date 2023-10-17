package sqlitedialect

type options struct {
	local bool
}

type Option func(o *options)

func applyOptions(ops ...Option) *options {
	opt := new(options)
	for _, o := range ops {
		o(opt)
	}
	return opt
}

// Local uses local timezone
func Local() Option {
	return func(o *options) {
		o.local = true
	}
}
