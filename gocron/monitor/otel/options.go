package otel

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	JobTimingAttribute = attribute.Key("gocron.job.timing")
)

type options struct {
	traceProvider trace.TracerProvider
	tracer        trace.Tracer //nolint:structcheck

	metricProvider metric.MeterProvider
	meter          metric.Meter

	attrs []attribute.KeyValue
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithAttributes(attrs ...attribute.KeyValue) Option {
	return func(o *options) {
		o.attrs = attrs
	}
}

func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(o *options) {
		o.traceProvider = provider
	}
}

func WithMeterProvider(provider metric.MeterProvider) Option {
	return func(o *options) {
		o.metricProvider = provider
	}
}
