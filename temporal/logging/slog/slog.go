package slog

import (
	"log/slog"

	"github.com/blink-io/x/temporal"
)

// slog.Logger can be used  in temporal directly
var _ temporal.Logger = (*slog.Logger)(nil)
