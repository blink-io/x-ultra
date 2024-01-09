package dbr

import (
	"log/slog"

	"github.com/gocraft/dbr/v2"
)

var _ dbr.EventReceiver = (*timingEventReceiver)(nil)

type timingEventReceiver struct {
	logger *slog.Logger
}

func NewTimingEventReceiver() dbr.EventReceiver {
	r := &timingEventReceiver{
		logger: slog.Default(),
	}
	return r
}

func (n *timingEventReceiver) Event(eventName string) {
	n.logger.Info("invoke Event", slog.String("eventName", eventName))
}

func (n *timingEventReceiver) EventKv(eventName string, kvs map[string]string) {
	n.logger.Info("invoke EventKv",
		slog.String("eventName", eventName),
		slog.Any("kvs", kvs),
	)
}

func (n *timingEventReceiver) EventErr(eventName string, err error) error {
	var errInfo string
	if err != nil {
		errInfo = err.Error()
	}
	n.logger.Info("invoke EventErr",
		slog.String("eventName", eventName),
		slog.Any("error", errInfo),
	)
	return err
}

func (n *timingEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	var errInfo string
	if err != nil {
		errInfo = err.Error()
	}
	n.logger.Info("invoke EventErrKv",
		slog.String("eventName", eventName),
		slog.Any("error", errInfo),
		slog.Any("kvs", kvs),
	)
	return err
}

func (n *timingEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	n.logger.Info("invoke EventErrKv",
		slog.String("eventName", eventName),
		slog.Int64("timingInNS", nanoseconds),
		slog.Any("kvs", kvs),
	)
}

// Timing receives the time an event took to happen.
func (n *timingEventReceiver) Timing(eventName string, nanoseconds int64) {
	slog.Default().Info("invoke Timing",
		slog.String("eventName", eventName),
		slog.Int64("timingInNS", nanoseconds),
	)
}
