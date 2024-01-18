package otel

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	instrumName = "github.com/blink-io/x/gocron/v2"
)

type monitor struct {
	opts *options

	ctx context.Context

	sc metric.Int64Counter

	fc metric.Int64Counter

	th metric.Float64Histogram
}

func NewMonitor(ops ...Option) (gocron.Monitor, error) {
	opts := applyOptions(ops...)
	if opts.traceProvider == nil {
		opts.traceProvider = otel.GetTracerProvider()
	}
	if opts.metricProvider == nil {
		opts.metricProvider = otel.GetMeterProvider()
	}
	if opts.meter == nil {
		opts.meter = opts.metricProvider.Meter(
			instrumName,
			metric.WithInstrumentationAttributes(opts.attrs...),
		)
	}
	if opts.tracer == nil {
		opts.tracer = opts.traceProvider.Tracer(
			instrumName,
			trace.WithInstrumentationAttributes(opts.attrs...),
		)
	}
	sc, err := opts.meter.Int64Counter("gocron.job.increment.success.count")
	if err != nil {
		otel.Handle(err)
		return nil, err
	}
	fc, err := opts.meter.Int64Counter("gocron.job.increment.failure.count")
	if err != nil {
		otel.Handle(err)
		return nil, err
	}
	th, err := opts.meter.Float64Histogram("gocron.job.timing.record")
	if err != nil {
		otel.Handle(err)
		return nil, err
	}
	m := &monitor{
		ctx: context.Background(),
		sc:  sc,
		fc:  fc,
		th:  th,
	}
	return m, nil
}

func (m *monitor) IncrementJob(id uuid.UUID, name string, tags []string, status gocron.JobStatus) {
	if gocron.Success == status {
		m.sc.Add(m.ctx, 1)
	} else if gocron.Fail == status {
		m.fc.Add(m.ctx, 1)
	}
}

func (m *monitor) RecordJobTiming(startTime, endTime time.Time, id uuid.UUID, name string, tags []string) {
	d := endTime.Sub(startTime)
	m.th.Record(m.ctx, milliseconds(d))

}

func milliseconds(d time.Duration) float64 {
	return float64(d) / float64(time.Millisecond)
}

var _ gocron.Monitor = (*monitor)(nil)
