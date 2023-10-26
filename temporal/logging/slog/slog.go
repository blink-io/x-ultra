package slog

import (
	"log/slog"

	"github.com/blink-io/x/temporal"
)

var _ temporal.Logger = (*slog.Logger)(nil)
