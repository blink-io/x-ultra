package dbr

import (
	"testing"
)

func TestTiming(t *testing.T) {
	er := NewTimingEventReceiver()
	er.Timing("eeee", 123)
}
