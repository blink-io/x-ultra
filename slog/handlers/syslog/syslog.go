package syslog

import (
	"log/slog"

	"github.com/samber/slog-syslog"
)

type Option = slogsyslog.Option

type Handler = slogsyslog.SyslogHandler

func NewHandler(o Option) slog.Handler {
	return o.NewSyslogHandler()
}
