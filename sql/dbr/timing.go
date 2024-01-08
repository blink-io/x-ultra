package dbr

import (
	"log/slog"

	"github.com/gocraft/dbr/v2"
)

type TimingEventReceiver struct {
	*dbr.NullEventReceiver
}

func NewTimingEventReceiver() dbr.EventReceiver {
	r := &TimingEventReceiver{
		NullEventReceiver: new(dbr.NullEventReceiver),
	}
	return r
}

// Timing receives the time an event took to happen.
func (n *TimingEventReceiver) Timing(eventName string, nanoseconds int64) {
	slog.Default().Info("DBR timing",
		slog.String("eventName", eventName),
		slog.Int64("costInNS", nanoseconds),
	)
}
