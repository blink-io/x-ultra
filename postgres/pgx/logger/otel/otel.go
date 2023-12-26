package otel

import (
	"github.com/exaring/otelpgx"
)

type (
	SpanNameFunc = otelpgx.SpanNameFunc

	Tracer = otelpgx.Tracer
)

var NewTracer = otelpgx.NewTracer
