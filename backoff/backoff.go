package backoff

import (
	"github.com/cenkalti/backoff/v4"
)

type BackOff = backoff.BackOff

var Retry = backoff.Retry
