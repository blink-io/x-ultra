package dbm

type options struct {
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

//
//func WithEventReceiver(er dbr.EventReceiver) Option {
//	return func(o *options) {
//		o.er = er
//	}
//}
