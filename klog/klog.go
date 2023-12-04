package klog

import (
	"github.com/go-kit/log"
)

type Logger = log.Logger

func Noop() Logger {
	return log.NewNopLogger()
}
