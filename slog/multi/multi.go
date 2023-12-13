package multi

import (
	"github.com/samber/slog-multi"
)

var (
	Fanout   = slogmulti.Fanout
	Failover = slogmulti.Failover
	Pipe     = slogmulti.Pipe
	Pool     = slogmulti.Pool
	Router   = slogmulti.Router
)

type (
	Middleware = slogmulti.Middleware
)
