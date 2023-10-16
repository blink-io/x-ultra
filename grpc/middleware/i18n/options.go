package i18n

import (
	"github.com/blink-io/x/i18n"
)

type option struct {
	header string
}

type Option func(*option)

func initOption(ops ...Option) *option {
	opt := &option{
		header: i18n.DefaultHeader,
	}

	for _, f := range ops {
		f(opt)
	}
	return opt
}

func Header(header string) Option {
	return func(o *option) {
		o.header = header
	}
}
