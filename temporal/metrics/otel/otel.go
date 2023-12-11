package otel

import (
	"time"

	"go.temporal.io/sdk/client"
)

type Handler struct {
}

var _ client.MetricsHandler = (*Handler)(nil)
var _ client.MetricsCounter = (*Handler)(nil)
var _ client.MetricsGauge = (*Handler)(nil)
var _ client.MetricsTimer = (*Handler)(nil)

func (h *Handler) WithTags(tags map[string]string) client.MetricsHandler {
	return h
}

func (h *Handler) Counter(name string) client.MetricsCounter {
	return h
}

func (h *Handler) Gauge(name string) client.MetricsGauge {
	return h
}

func (h *Handler) Timer(name string) client.MetricsTimer {
	return h
}

func (h *Handler) Inc(i int64) {
}

func (h *Handler) Record(duration time.Duration) {
}

func (h *Handler) Update(gauge float64) {
}
