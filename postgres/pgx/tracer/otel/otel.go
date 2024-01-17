package otel

import (
	"github.com/exaring/otelpgx"
)

type (
	SpanNameFunc = otelpgx.SpanNameFunc

	Tracer = otelpgx.Tracer

	Option = otelpgx.Option
)

var NewTracer = otelpgx.NewTracer
