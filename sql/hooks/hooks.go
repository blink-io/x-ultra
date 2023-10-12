package hooks

import (
	"github.com/qustavo/sqlhooks/v2"
)

type Hooks = sqlhooks.Hooks

var (
	Wrap    = sqlhooks.Wrap
	Compose = sqlhooks.Compose
)
