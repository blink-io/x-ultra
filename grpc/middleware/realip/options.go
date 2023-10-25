package realip

import (
	"github.com/blink-io/x/realip"
)

type options = realip.Options

type Option func(*options)

func initOption(ops ...Option) *options {
	opt := realip.DefaultOptions
	for _, o := range ops {
		o(opt)
	}
	return opt
}
